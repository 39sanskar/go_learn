package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string // you have to mention explictly hoe much data have to come in, here that is 4

	fruitList[0] = "Apple"  // at the 0th position add "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("FruitList is: ", fruitList) // FruitList is:  [Apple Tomato  Peach]
	fmt.Println("FruitList is: ", len(fruitList)) // FruitList is:  4


	var vegList = [3]string{"potato", "beans", "mushroom"} // kind of initialise it and put some value at the same time.
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list is: ", len(vegList))

}