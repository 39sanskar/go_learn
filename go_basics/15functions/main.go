package main

import "fmt"

func main(){
	fmt.Println("Welcome to functions in golang")
	greeter()

	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
  fmt.Println("Pro Message is: ", myMessage)
  
	/*
	// Not allow function inside function
	func greeterTwo(){
		fmt.Println("Another method")
	}
	greeterTwo()
	*/

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string){
	total := 0

	for _, val := range values{
		total += val 
	}

	return total, "Hi Pro result function"
}

func greeterTwo(){
	fmt.Println("Another method")
}

func greeter(){
	fmt.Println("Namastey from golang")
}

// func proAdder(values ...int) => means you’re declaring a function named proAdder that accepts a variadic parameter of type int.

// A variadic parameter allows the caller to pass zero or more arguments of that type. Inside the function, values behaves like a slice of int.
