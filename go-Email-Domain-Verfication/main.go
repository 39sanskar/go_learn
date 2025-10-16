package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	DefaultDNSTimeout    = 5 * time.Second // DefaultDNSTimeout specifies the timeout for DNS lookups.
	MaxConcurrentDomains = 10              // MaxConcurrentDomains limits the number of domains/emails processed concurrently.
	OutputFormatJSON     = "json"          // OutputFormatJSON is for JSON output.
	OutputFormatCSV      = "csv"           // OutputFormatCSV is for CSV output.
	SMTPVerificationRetry = 2              // number of attempts for SMPT verification
)

// EmailVerificationReport holds verification details for an email
type EmailVerificationReport struct {
	Email            string `json:"email"`
	ValidFormat      bool   `json:"validFormat"`
	ValidDomain      bool   `json:"validDomain"`
	DomainHasMX      bool   `json:"domainHasMX"`
	DomainSPF        string `json:"domainSPF,omitempty"`
	DomainSPFError   string `json:"domainSPFError,omitempty"`
	DomainDMARC      string `json:"domainDMARC,omitempty"`
	DomainDMARCError string `json:"domainDMARCError,omitempty"`
	MailboxExists    bool   `json:"mailboxExists"`
	Error            string `json:"error,omitempty"`
}

// DomainVerificationReport holds verification details for a domain
type DomainVerificationReport struct {
	Domain        string   `json:"domain"`
	HasMX         bool     `json:"hasMx"`
	MXRecords     []string `json:"mxRecords,omitempty"`
	HasSPF        bool     `json:"hasSpf"`
	SPFRecord     string   `json:"spfRecord,omitempty"`
	SPFError      string   `json:"spfError,omitempty"`
	HasDMARC      bool     `json:"hasDmarc"`
	DMARCRecord   string   `json:"dmarcRecord,omitempty"`
	DMARCError    string   `json:"dmarcError,omitempty"`
	DNSSecEnabled bool     `json:"dnsSecEnabled"` // Placeholder for DNSSEC
	Error         string   `json:"error,omitempty"`
}

// ResultWrapper to hold either an Email or Domain report, allowing a single channel for results
type ResultWrapper struct {
	EmailReport  *EmailVerificationReport  `json:"emailReport,omitempty"`
	DomainReport *DomainVerificationReport `json:"domainReport,omitempty"`
}

// Config for the program application.
type Config struct {
	Timeout       time.Duration
	ConcurrentOps int
	OutputFormat  string
	Verbose       bool
	InputType     string // "email", "domain", or "auto"
	VerifySMTP    bool
}

func main() {
	config := parseConfig()

	// Initialize DNS resolver with context and timeout
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: config.Timeout,
			}
			// Using Google's public DNS for consistency
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	items := make(chan string)
	results := make(chan ResultWrapper)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < config.ConcurrentOps; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range items {
				itemType := config.InputType

				// Auto-detect type if needed
				if itemType == "auto" {
					if isValidEmailFormat(item) {
						itemType = "email"
						if config.Verbose {
							log.Printf("Detected email: %s", item)
						}
					} else {
						itemType = "domain"
						if config.Verbose {
							log.Printf("Detected domain: %s", item)
						}
					}
				}

				if itemType == "email" {
					r := verifyEmail(item, resolver, config)
					results <- ResultWrapper{EmailReport: &r}
				} else {
					r := verifyDomain(item, resolver, config)
					results <- ResultWrapper{DomainReport: &r}
				}
			}
		}()
	}

	// Read from stdin
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if config.Verbose {
			log.Println("Reading input (emails/domains)...")
		}
		for scanner.Scan() {
			item := strings.TrimSpace(scanner.Text())
			if item != "" {
				items <- item
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		close(items)
	}()

	// Collect and print reports
	go func() {
		wg.Wait() // Wait for all workers to finish
		close(results)
	}()

	for res := range results {
		printUnifiedReport(res, config.OutputFormat)
	}

	if config.Verbose {
		log.Println("Verification complete.")
	}
}

// parseConfig handles command-line arguments and sets up the application configuration.
func parseConfig() Config {
	config := Config{
		Timeout:       DefaultDNSTimeout,
		ConcurrentOps: MaxConcurrentDomains,
		OutputFormat:  OutputFormatCSV, // Default output
		InputType:     "auto",
		Verbose:       false,
		VerifySMTP:    false,
	}

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-t", "--timeout":
			if i+1 < len(args) {
				duration, err := time.ParseDuration(args[i+1])
				if err != nil {
					log.Fatalf("Invalid timeout duration: %s. Use format like '5s' or '1m'.", args[i+1])
				}
				config.Timeout = duration
				i++
			}
		case "-c", "--concurrent":
			if i+1 < len(args) {
				count, err := fmt.Sscanf(args[i+1], "%d", &config.ConcurrentOps)
				if err != nil || count != 1 || config.ConcurrentOps <= 0 {
					log.Fatalf("Invalid concurrent operations count: %s. Must be a positive integer.", args[i+1])
				}
				i++
			}
		case "-o", "--output":
			if i+1 < len(args) {
				format := strings.ToLower(args[i+1])
				if format != OutputFormatCSV && format != OutputFormatJSON {
					log.Fatalf("Invalid output format: %s. Supported formats are 'csv' and 'json'.", args[i+1])
				}
				config.OutputFormat = format
				i++
			}
		case "-v", "--verbose":
			config.Verbose = true
		case "-T", "--type":
			if i+1 < len(args) {
				t := strings.ToLower(args[i+1])
				if t != "email" && t != "domain" && t != "auto" {
					log.Fatalf("Invalid type: %s", t)
				}
				config.InputType = t
				i++
			} 	
		case "--smtp":
			config.VerifySMTP = true
		case "-h", "--help":
			fmt.Printf(`Usage: go run main.go [options]

Options:
  -T, --type <email|domain|auto>  Input type (default: auto)
  -t, --timeout <duration>   DNS lookup timeout (default 5s)
  -c, --concurrent <count>   Max concurrent verifications (default 10)
  -o, --output <csv|json>    Output format (default csv)
  -v, --verbose              Enable verbose logging
  --smtp                     Attempt SMTP mailbox verification (only for emails)
  -h, --help                 Show this message

Input: Provide emails or domains line by line via stdin.
Example:
  echo "example.com" | go run main.go
  echo "user@gmail.com" | go run main.go -o json
  cat list.txt | go run main.go -T auto --smtp -v
`)
			os.Exit(0)
		}
	}
	return config
}

// verifyEmail performs full email verification
func verifyEmail(email string, resolver *net.Resolver, config Config) EmailVerificationReport {
	report := EmailVerificationReport{Email: email}

	if config.Verbose {
		log.Printf("Verifying email: %s", email)
	}

	// Step 1: Email format
	if !isValidEmailFormat(email) {
		report.ValidFormat = false
		report.Error = "Invalid email format"
		return report
	}
	report.ValidFormat = true

	// Step 2: Extract domain
	parts := strings.Split(email, "@")
	domain := parts[1]

	// Step 3: Domain format (basic check)
	if !isValidDomainFormat(domain) {
		report.ValidDomain = false
		report.Error = "Invalid domain format derived from email"
		return report
	}
	report.ValidDomain = true

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	// Step 4: MX check
	mxRecords, err := resolver.LookupMX(ctx, domain)
	if err != nil {
		if report.Error == "" { // Don't overwrite more critical errors
			report.Error = fmt.Sprintf("MX lookup failed for email domain: %v", err)
		}
	} else if len(mxRecords) > 0 {
		report.DomainHasMX = true
	}

	// Step 5: SPF check (simplified)
	txtRecords, err := resolver.LookupTXT(ctx, domain)
	if err == nil {
		for _, r := range txtRecords {
			if strings.HasPrefix(r, "v=spf1") {
				report.DomainSPF = r
				if !isValidSPF(r) { // Basic syntax check
					report.DomainSPFError = "Potentially invalid SPF syntax"
				}
				break
			}
		}
	} else if config.Verbose {
		log.Printf("Could not lookup TXT for %s (SPF check): %v", domain, err)
	}

	// Step 6: DMARC check (simplified)
	dmarcRecords, err := resolver.LookupTXT(ctx, "_dmarc."+domain)
	if err == nil {
		for _, r := range dmarcRecords {
			if strings.HasPrefix(r, "v=DMARC1") {
				report.DomainDMARC = r
				if !isValidDMARC(r) { // Basic syntax check
					report.DomainDMARCError = "Potentially invalid DMARC syntax"
				}
				break
			}
		}
	} else if config.Verbose {
		log.Printf("Could not lookup TXT for _dmarc.%s (DMARC check): %v", domain, err)
	}

	// Step 7: SMTP mailbox verification (optional)
	if config.VerifySMTP && report.DomainHasMX && len(mxRecords) > 0 {
		// Try all MX records until one succeeds
		mailboxFound := false
		for _, mx := range mxRecords {
			mxHost := strings.TrimSuffix(mx.Host, ".")
			if verifySMTP(email, mxHost, config.Verbose) {
				mailboxFound = true
				break
			}
		}
		report.MailboxExists = mailboxFound
	}

	return report
}

// verifyDomain performs DNS record checks for a domain. This is a refactored checkDomain.
func verifyDomain(domain string, resolver *net.Resolver, config Config) DomainVerificationReport {
	report := DomainVerificationReport{
		Domain: domain,
	}
	if config.Verbose {
		log.Printf("Verifying domain: %s", domain)
	}

	// Basic domain format validation
	if !isValidDomainFormat(domain) {
		report.Error = "Invalid domain format"
		return report
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	// MX Records
	mxRecords, err := resolver.LookupMX(ctx, domain)
	if err != nil {
		report.Error = fmt.Sprintf("MX lookup failed: %v", err)
		if config.Verbose {
			log.Printf("Error for %s (MX): %v", domain, err)
		}
	} else if len(mxRecords) > 0 {
		report.HasMX = true
		for _, mx := range mxRecords {
			report.MXRecords = append(report.MXRecords, mx.Host)
		}
	}

	// SPF Records (TXT records)
	txtRecords, err := resolver.LookupTXT(ctx, domain)
	if err != nil {
		if config.Verbose {
			log.Printf("Error for %s (TXT): %v", domain, err)
		}
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				report.HasSPF = true
				report.SPFRecord = record
				// Basic SPF validation (more advanced validation can be added)
				if !isValidSPF(record) {
					report.SPFError = "Potentially invalid SPF record syntax"
				}
				break
			}
		}
	}

	// DMARC Records (TXT records for _dmarc.domain)
	dmarcDomain := "_dmarc." + domain
	dmarcRecords, err := resolver.LookupTXT(ctx, dmarcDomain)
	if err != nil {
		if config.Verbose {
			log.Printf("Error for %s (DMARC): %v", domain, err)
		}
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				report.HasDMARC = true
				report.DMARCRecord = record
				// Basic DMARC validation
				if !isValidDMARC(record) {
					report.DMARCError = "Potentially invalid DMARC record syntax"
				}
				break
			}
		}
	}

	// DNSSEC Check (placeholder - actual implementation is complex)
	// report.DNSSecEnabled = checkDNSSec(resolver, domain, config.Verbose)

	return report
}

// ---- Validation helpers ---

// Basic email regex
func isValidEmailFormat(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func isValidDomainFormat(domain string) bool {
	if len(domain) == 0 || len(domain) > 253 {
		return false
	}
	if !strings.Contains(domain, ".") {
		return false
	}
	validDomainRegex := regexp.MustCompile(`^(?i)[a-z0-9-\.]+$`)
	if !validDomainRegex.MatchString(domain) {
		return false
	}

	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if len(label) == 0 || len(label) > 63 {
			return false
		}
		if strings.HasPrefix(label, "-") || strings.HasSuffix(label, "-") {
			return false
		}
	}
	return true
}

// Basic SPF check
func isValidSPF(spf string) bool {
	return strings.HasPrefix(spf, "v=spf1") // Can be enhanced with more regex or parsing
}

// Basic DMARC check
func isValidDMARC(dmarc string) bool {
	return strings.HasPrefix(dmarc, "v=DMARC1") && strings.Contains(dmarc, "p=") // Can be enhanced
}

// Simple SMTP RCPT TO verification
func verifySMTP(email, mxHost string, verbose bool) bool {
	for i := 0; i < SMTPVerificationRetry; i++ {
		if verbose {
			log.Printf("Attempting SMTP verification for %s via MX %s (attempt %d/%d)", email, mxHost, i+1, SMTPVerificationRetry)
		}
		client, err := smtp.Dial(fmt.Sprintf("%s:25", mxHost))
		if err != nil {
			if verbose {
				log.Printf("SMTP Dial failed for %s to %s: %v", email, mxHost, err)
			}
			continue
		}
		defer client.Close()

		if err := client.Hello("example.com"); err != nil {
			if verbose {
				log.Printf("SMTP HELO failed for %s to %s: %v", email, mxHost, err)
			}
			continue
		}
		if err := client.Mail("check@example.com"); err != nil { // Sender address for the check
			if verbose {
				log.Printf("SMTP MAIL FROM failed for %s to %s: %v", email, mxHost, err)
			}
			continue
		}
		if err := client.Rcpt(email); err == nil {
			if verbose {
				log.Printf("SMTP RCPT TO for %s succeeded via %s", email, mxHost)
			}
			return true
		} else {
			if verbose {
				log.Printf("SMTP RCPT TO for %s failed via %s: %v", email, mxHost, err)
			}
		}
	}
	return false
}

// printUnifiedReport prints a single verification report in the specified format, handling both types.
func printUnifiedReport(res ResultWrapper, format string) {
	switch format {
	case OutputFormatCSV:
		
		if res.EmailReport != nil {
			fmt.Printf("email,validFormat,validDomain,domainHasMX,domainSPF,domainSPFError,domainDMARC,domainDMARCError,mailboxExists,error\n")
			r := res.EmailReport
			fmt.Printf("%s,%t,%t,%t,\"%s\",\"%s\",\"%s\",\"%s\",%t,\"%s\"\n",
				r.Email,
				r.ValidFormat,
				r.ValidDomain,
				r.DomainHasMX,
				r.DomainSPF,
				r.DomainSPFError,
				r.DomainDMARC,
				r.DomainDMARCError,
				r.MailboxExists,
				r.Error,
			)
		} else if res.DomainReport != nil {
			fmt.Printf("domain,hasMx,mxRecords,hasSpf,spfRecord,spfError,hasDmarc,dmarcRecord,dmarcError,dnsSecEnabled,error\n")
			r := res.DomainReport
			mxRecords := strings.Join(r.MXRecords, ";")
			fmt.Printf("%s,%t,\"%s\",%t,\"%s\",\"%s\",%t,\"%s\",\"%s\",%t,\"%s\"\n",
				r.Domain,
				r.HasMX,
				mxRecords,
				r.HasSPF,
				r.SPFRecord,
				r.SPFError,
				r.HasDMARC,
				r.DMARCRecord,
				r.DMARCError,
				r.DNSSecEnabled,
				r.Error,
			)
		}
	case OutputFormatJSON:
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ") // Pretty print JSON
		if err := encoder.Encode(res); err != nil {
			log.Printf("Error encoding JSON report: %v", err)
		}
	}
}
