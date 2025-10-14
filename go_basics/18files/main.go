package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Coding Hub."

	file, err := os.Create("./mygofile.txt")
	
	// checking of the error
	// if err != nil {
	// 	  panic(err) // panic will shutdown the execution of the program and will show you what the error you are facing.
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	// if it has executed it gives you back length or it might give you error.

	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
  readFile("./mygofile.txt")
}

func readFile(filename string){
	 databyte, err :=  ioutil.ReadFile(filename)
	 // whenever you read the file its not being read into the string format specially ehen you read data from internet as well always read in the bytes format.

	 checkNilErr(err)

	 fmt.Println("Text data inside the file is \n", databyte)

}

// creating a function checkNilErr bec many times err syntax is repeating.
func checkNilErr(err error){
	 if err != nil {
		  panic(err)
	 }
}