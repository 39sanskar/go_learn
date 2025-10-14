// If you want Go to continue into the next case, you explicitly write fallthrough.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Dice game with fallthrough")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can open the game")
	case 6:
		fmt.Println("You rolled a 6, move 6 spots")
		fallthrough // force execution of the next case
	case 5:
		fmt.Println("Bonus: You get an extra turn!")
	default:
		fmt.Printf("You can move %d spots\n", diceNumber)
	}
}

/*

What %d means
- %d → format an integer in base 10 (decimal).
- It’s commonly used for int, int32, int64, etc.

Other related verbs
%d → decimal (base 10)
%b → binary
%o → octal
%x / %X → hexadecimal (lower/upper case)
%c → character (Unicode code point)
%q → quoted character/string

*/

/*

- Important rules about fallthrough
- It always goes only to the very next case (not multiple).
- It does not re-check the condition of the next case — it just executes it blindly.
- Usually used rarely, when you want grouped behavior but still need some case-specific code.

*/