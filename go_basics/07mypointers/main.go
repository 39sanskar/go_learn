package main

import "fmt"
func main() {
	fmt.Println("Welcome to a class on pointers")

	/*

	// Declaration of the pointer

	var ptr *int 

	fmt.Println("Value of pointer is ", ptr)
	// if you start a pointer you dont put any value to it, its value is actually nil (default value is nil)

 */

  
 myNumber := 23

 var myptr = &myNumber  // creating a pointer which is also rferencing to some memory 
 // whenever their is a talk about referencing thats where you use  &

 // Pointer is reference to the direct memory location
 fmt.Println("Value of actual pointer is ", myptr) // 0xc0000100f8  // provide the acctual memory location
 fmt.Println("Value of actual pointer is ", *myptr)  // 23   // provide the acctual value

 // when you say  *myptr that means hey i want to see what inside that pointer

 *myptr = *myptr + 2
 fmt.Println("New value is: ", myNumber) // 25
 
}