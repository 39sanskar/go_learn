package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // in the documentation it is mention, we are always use.
	// this is standard for formatting.

	// if you want to print current day
	fmt.Println(presentTime.Format("01-02-2006 Monday")) // always have to give Monday here.

	// print current time date month and day (Syntax)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// May be you want to create the time from some of the value you want to manually entry, you need to go other way around as well.
	createdDate := time.Date(2020, time.September, 12, 23, 23, 0, 0,  time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
} 
