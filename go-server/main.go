package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Home page (static files served by http.FileServer)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Hello! hi How are you, friend?")
}

// Handle form submissions (POST)
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	if name == "" || address == "" {
		http.Error(w, "Both name and address are required!", http.StatusBadRequest)
		return
	}

	// Use a small HTML template for a nice response
	tmpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head><meta charset="UTF-8"><title>Form Submitted</title></head>
		<body style="font-family: Arial; text-align: center; margin-top: 50px;">
			<h2>Form Submitted Successfully!</h2>
			<p><b>Name:</b> {{.Name}}</p>
			<p><b>Address:</b> {{.Address}}</p>
			<a href="/">Go Back</a>
		</body>
		</html>
	`

	t := template.Must(template.New("response").Parse(tmpl))
	data := struct {
		Name, Address string
	}{
		Name:    name,
		Address: address,
	}
	t.Execute(w, data)
}

// Simple health check route
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","time":"%s"}`, time.Now().Format(time.RFC3339))
}

func main() {
	// Serve static files from the "static" folder
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Routes
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Starting server on http://localhost:7070")
	log.Fatal(http.ListenAndServe(":7070", logRequest(http.DefaultServeMux)))
}

// Middleware for logging requests
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("%s %s [%s]", r.Method, r.URL.Path, time.Since(start))
	})
}
