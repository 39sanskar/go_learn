package main 

import "fmt"

func main(){
	fmt.Println("Structs in golang")
	// structs is the alternative of the classes
	// no inheritance in golang; No super or parent

	sanskar := User{"Sanskar", "sanskar@go.dev", true, 21}
	fmt.Println(sanskar)
	fmt.Printf("sanskar detaild are: %+v\n", sanskar)
	fmt.Printf("Name is %v and email is %v.", sanskar.Name, sanskar.Email)
	
}

// using %+v in the case of structure
type User struct {
	Name string 
	Email string
	Status bool
	Age int
}