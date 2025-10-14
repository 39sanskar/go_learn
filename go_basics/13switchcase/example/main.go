package main

import "fmt"

func main() {
    n := 65
    fmt.Printf("Decimal: %d\n", n)
    fmt.Printf("Binary: %b\n", n)
    fmt.Printf("Octal: %o\n", n)
    fmt.Printf("Hex: %x\n", n)
    fmt.Printf("Char: %c\n", n)
}

/*

Output:
Decimal: 65
Binary: 1000001
Octal: 101
Hex: 41
Char: A

*/


/*

Other related verbs
%d → decimal (base 10)
%b → binary
%o → octal
%x / %X → hexadecimal (lower/upper case)
%c → character (Unicode code point)
%q → quoted character/string

*/