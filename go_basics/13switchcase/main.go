package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
  diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
  case 2:
		fmt.Println("You can move 2 spot")	
	case 3:
		fmt.Println("You can move 3 spot")
  case 4: 
	  fmt.Println("You can move 4 spot")	
	case 5:
		fmt.Println("You can move 5 spot")	
	case 6:
		fmt.Println("You can move to 6 spot and roll the dice again")		
	default:
		fmt.Println("What was that!")			
	}
}

/*

- rand.Seed is deprecated: As of Go 1.20 there is no reason to call Seed with
a random value. Programs that call Seed with a known value to get
a specific sequence of results should use New(NewSource(seed)) to
obtain a local random generator

- But in Go 1.20+
- The global generator is already seeded automatically with a secure random value.
- That means calling rand.Seed(time.Now().UnixNano()) is no longer necessary (and Go marks it as deprecated).

*/

/*

- What’s happening

- rand.Seed(time.Now().UnixNano())
- Seeds the random number generator with the current time in nanoseconds.
- Ensures you don’t get the same dice value every time you run the program.

- rand.Intn(6) + 1
- Generates a random integer from 0–5 and adds 1.
- Final result = 1–6 (like a real dice).

- switch diceNumber
- Matches the random dice value with one of the case branches.
- default is like the "else" — but in this case, it won’t ever run since diceNumber is always between 1–6.

*/