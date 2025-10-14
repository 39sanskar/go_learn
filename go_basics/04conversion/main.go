package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	fmt.Println("Welcome to our work app")
	fmt.Println("Please rate our Work in between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // keep you reading till you hit a  \n (Newline)

	fmt.Println("Thanks for rating, ", input) // input is string type, so we cant add  input + 1
  
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // here handle err, and also handle string and convert it into number.

	if err != nil {
		fmt.Println(err)
	  
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}


/*

strings.TrimSpace(input)

🔹 What it does
- The function strings.TrimSpace (from Go’s strings package) removes all the leading and trailing whitespace characters from a string.
Whitespace includes:
- Space " "
- Tab \t
- Newline \n
- Carriage return \r

*/