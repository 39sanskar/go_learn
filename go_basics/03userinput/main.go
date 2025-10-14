package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Work:")
	
	// comma ok syntax || err err syntax 

	input, _ := reader.ReadString('\n') // keep reading as soon as \n is there 
	fmt.Println("Thanks for rating, ",  input)
	fmt.Printf("Type of this rating is %T", input) // Type of this rating is string

	
	// in language go we don't have try catch if something goes wrong their is nobody to catch that, the language design aspect you treat problems or error is something like variable , like true or false value this is was exactly comma ok syntax is.

	// short variable declaration := sometimes when you use packages like, bufio  etc.  you dont know what is comming up that exactly this short variable declaration :=  is designed.
}


// pkg.go.dev it is official website for Go packages.

/*
bufio Package in Go
- The bufio package in Go provides buffered input and output.
- It wraps around an io.Reader or io.Writer (like files, stdin, sockets, etc.) to make reading/writing more efficient and convenient.

🔹 Why use bufio?
- Without buffering → reading/writing happens one small chunk at a time → slower.
- With bufio → data is stored in memory temporarily → fewer system calls → faster performance.


🔹 Commonly Used Types & Functions

- bufio.NewReader(r io.Reader)
Creates a buffered reader.

- bufio.NewWriter(w io.Writer)
Creates a buffered writer.

- Reader methods:
- ReadString(delim byte) → read until a delimiter (like '\n').
- ReadBytes(delim byte) → same but returns []byte.
- ReadLine() → reads a line (without newline).
- Peek(n int) → look at next n bytes without advancing.

Writer methods:
- Write(p []byte) → write data.
- WriteString(s string) → write string.
- Flush() → force buffered data to be written out.

✅ Summary:
- bufio is for efficient buffered I/O.
- Use NewReader for input, NewWriter for output.
- Always Flush() writers to push data out.

*/