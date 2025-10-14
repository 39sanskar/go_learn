package main 

import "fmt"

func main(){
	fmt.Println("Structs in golang")
	// structs is the alternative of the classes
	// no inheritance in golang; No super or parent

	sanskar := User{"Sanskar", "sanskar@go.dev", true, 21}
	fmt.Println(sanskar)
	fmt.Printf("sanskar detaild are: %+v\n", sanskar)
	fmt.Printf("Name is %v and email is:  %v\n", sanskar.Name, sanskar.Email)
	sanskar.GetStatus()
	sanskar.NewMail()
	fmt.Printf("Name is %v and email is:  %v\n", sanskar.Name, sanskar.Email)
  // whenever you pass on these objects or the structs it actually passes on a copy.

	// if you want to really pass on the original object you should be passing up the reference of it or a Pointer to that.
}

// using %+v in the case of structure
type User struct {
	Name string 
	Email string
	Status bool
	Age int 
}

func (u User) GetStatus(){
	// if you are planning to export this method make sure first letter is capital in GetStatus.
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	// if you are planning to export this method make sure first letter is capital in NewMail.
	u.Email = "program@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

