package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23  // arbitiary value 
  var result string

	if loginCount < 10 { // here not allow to move curley bracket in the next line.
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		 result = "Exactly 10 login count"
  }

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// cases of web request handling 

	if num := 3; num < 10 {
		fmt.Println("Num is not less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// if err != nil {

	// }

}

/*

Thats why this works:

if loginCount < 10 {
  result = "Regular user"
}

But this ❌ will give a syntax error:

if loginCount < 10
{
  result = "Regular user"
}

- Go enforces this because of the automatic semicolon insertion rule in the compiler.
- When you press Enter after if loginCount < 10, Go thinks the statement ends there and automatically inserts a semicolon ;, which breaks the syntax.

*/