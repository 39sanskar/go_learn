package main

import (
	//"math/rand"
	"crypto/rand"
	"math/big"
	"fmt"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 12
	// var mynumberTwo float64 = 8.6

	// fmt.Println("The sum is: ", mynumberOne + int(mynumberTwo))

	// random number from math

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)


	// generating random number by cryptography (govern by crypto it is much more accurate)
  
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
  
}