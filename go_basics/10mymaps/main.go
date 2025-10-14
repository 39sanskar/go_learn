package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	// Create an empty map using make
	languages := make(map[string]string) // map[keyType]valueType


	// Add key-value pairs
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// Add key-value pairs
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// Print the entire map
	fmt.Println("List of all languages: ", languages)

	// Access specific keys
	fmt.Println("JS shorts for: ", languages["JS"])
	fmt.Println("PY shorts for: ", languages["PY"])

	// Delete a key
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// Re-add a key-value pair
	languages["RB"] = "Ruby"
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		  fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// sometimes you dont care about what the first operator so replace with _
	for _, value := range languages {
		  fmt.Printf("For key v, value is %v\n", value)
	}
}

// %v => it is for value 
// %T => it is for type