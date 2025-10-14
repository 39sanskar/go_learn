package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices ")

	// Creating a slice (no size mentioned like arrays).
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList) // []string

	// Append adds new elements to the slice.
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	// Output: [Apple Tomato Peach Mango Banana]

	// Slice from index 1 till the end.
	fruitList = append(fruitList[1:]) 
	fmt.Println(fruitList)
	// Output: [Tomato Peach Mango Banana]

	// Slice from index 1 (inclusive) to 3 (exclusive).
	fruitList = append(fruitList[1:3]) 
	fmt.Println(fruitList)
	// Output: [Peach Mango]

	fruitList = append(fruitList[:3]) // by-default started from 0 (inclusive)
	fmt.Println(fruitList) // [Peach Mango Banana]


	highScores := make([]int, 4) // using make to create a slices

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321) 
	// use method append it is going to reallocate the memory and all of the value accomodated.
	// When you fill up a slice and then append, Go reallocates a new backing array with bigger capacity (usually doubling).

	fmt.Println("Before sorting:", highScores) // [234 945 465 867 555 666 321]

  fmt.Println(sort.IntsAreSorted(highScores)) // false
	// method is avilable in the slices not in the array
	sort.Ints(highScores) // sorts slice in-place in ascending order
	fmt.Println("After sorting: ", highScores) // [234 321 465 555 666 867 945]
	fmt.Println(sort.IntsAreSorted(highScores))  // true

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2 
	// append is using add and remove the value.
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	
}

/*

Key Points:

-Slices vs Arrays
- Arrays: fixed length ([3]string{"a","b","c"})
- Slices: dynamic, no fixed length ([]string{"a","b","c"})

- Appending
append(slice, elements...) → creates a new slice with the added elements.

- Slicing
- slice[start:end] → includes start, excludes end.
- slice[1:] → from index 1 till end.
- slice[:3] → from start till index 3 (excluding 3).

Why append(fruitList[1:]) works
Normally, we’d just write fruitList = fruitList[1:].
Using append(fruitList[1:]) works too, because append with a single slice argument just copies it into a new slice.
(Though it’s not common — usually people just use fruitList = fruitList[1:].)

*/