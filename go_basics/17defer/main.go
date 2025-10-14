package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("World")

	// whatever you put the defer it just cut the line and place it just above the curley bracket. (in the reverse order they are defered)
	// LIFO => using last in first out principle.
	// special case in defer 
	defer fmt.Println("Duniya")
	fmt.Println("How are you dear friend?")

	defer fmt.Println("Aman")
	defer fmt.Println("Ajay")
  defer fmt.Println("Raj")
	defer fmt.Println("VidyaSagar")
	fmt.Println("Welcome")

	myDefer() 
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// when you see a defer keyword then execution happen a little bit later, a macdonalds queue is been created. it might be many macdonalds or just one macdonalds.