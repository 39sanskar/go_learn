package main

import (
	"fmt"
)

func main() {
	courses := []string{"ReactJS", "JavaScript", "Swift", "Python", "Ruby"}
	fmt.Println("Before:", courses)

	index := 2 // remove "Swift"
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("After: ", courses)

}

/*

Before: [ReactJS JavaScript Swift Python Ruby]
After:  [ReactJS JavaScript Python Ruby]

*/