package main

import "fmt"

const LoginToken string = "xyzankdfhiehiefbhfd"  // Public
/*
I am creating LoginToken  with firstcharacter as capital "L" , capital L has significance importance here reason for that is this is now public variable. 

in golang  declaring a public variable using first letter is capital , this LoginToken is accessible by any other files into this folder or actually in this  program. 
*/

// jwtToken := 3000000  // outside method short variable declaration operator that is not allow.

// var jwtToken = 300000  // this is allow, using var keyword

// var jwtToken int = 300000   // this is also allow, using var keyword
func main() {
	// to create a variable we use a keyword var
	var username string = "Aman"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)


	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)


	var smallFloat float32 = 255.45544511254451885
	fmt.Println(smallFloat) // 255.45544  // in the case of float32 i get 5 values after the decimal (this is valid when value is long.)
	fmt.Printf("Variable is of type: %T \n", smallFloat) // Variable is of type: float32 

  
	var bigFloat float64 = 255.45544511254451885
	fmt.Println(bigFloat) // 255.4554451125445   // in the case of float64 i get more precise value after decimal
	fmt.Printf("Variable is of type: %T \n", bigFloat)


	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable) // 0   // it does not any garbage value, it's always being expect if you declare an integer it is going to  be 0
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "learningcode.in" // the variable type should be declare, and it is very important for golang.
	// the lexer comes in if you are not going to say what type of variable it is i will decide it for you based for on what value putting it up. So that here the type of website is string  
	fmt.Println(website)
	// website = 3  // that is not allow type change

  // no var style => you can totally ignore the keyword var and still can declare other variable.


	// In Go, the := operator is called the short variable declaration operator.

	numberOfUser := 300000.0 // inside method (:= short variable declaration operator is allow ) but outside method that is not allow
	fmt.Println(numberOfUser)

	/*
	numberOfUser := 300000.0
  := is the short variable declaration in Go.
  - It both declares and initializes the variable.
  300000.0 is written with a decimal point → so Go treats it as a float64 by default (a floating-point number).
  Therefore, numberOfUser will have the type float64.
	*/


	fmt.Println(LoginToken) // xyzankdfhiehiefbhfd
	fmt.Printf("Variable is of type: %T \n", LoginToken) // string 
}


// short variable declaration := sometimes when you use packages like, bufio  etc.  you dont know what is comming up that exactly this short variable declaration :=  is designed.
