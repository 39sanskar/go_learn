package main

import "fmt"

func main(){
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// d< len(days) => this is little bit problematic in case you are using array, len(days) it will give you all the values in the arrays their might be chance value is present in the array or not. but in slices it automatically adjust that.
	// their is no such thing as ++d. always go for d++

	// for d := 0; d< len(days); d++{
	// 	fmt.Println(days[d])
	// }


	// here using keyword range this range automatically loops through over any array or a slice.

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }


	// for _, day := range days {
	// 	fmt.Printf("index is and value is %v\n", day)
	// }


	 
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco // use goto statement to move on to any label 
		}

		if rougueValue == 5 {
			rougueValue++
			continue // continue just skip over one phase
		}

		if rougueValue == 7 {
			rougueValue++
			break  // break just terminate the loops all together 
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at sanskar.dev.com")

}