package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`

	Price    int

	Platform string `json:"website"`

	Password string `json:"-"` // here dash "-" simply means i dont want this field to reflected  whoever is consuming my api. 

	Tags     []string `json:"tags,omitempty"` // omitempty is simply says that if the value is nil then dont throw that field at all. 
	// you have to very careful about the space here `json:"tags,omitempty"`
}

func main(){
	fmt.Println(" Welcome to JSON video ")
	// EncodeJson()
	DecodeJson()
}

// Encoding the JSON => this simply means i have a data it can be slices , arrays keyvalues pair whatever that is i want to convert the data into a valid json.

func EncodeJson(){

  CoadingCourses := []course{
		{"ReactJS Bootcamp", 299, "CoadingHub.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "CoadingHub.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "CoadingHub.in", "pqs123", nil},
	}

	// package this data as JSON data 

	// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
	// inside Marshal() pass on interface.

	finalJson, err := json.MarshalIndent(CoadingCourses, "", "\t") // it might show some error, here need error handling.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson(){
	  jsonDataFromWeb := []byte(`
		{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "CoadingHub.in",
		"tags": ["web-dev","js"]
	}	
	`)

	var CoadingCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &CoadingCourse)
		fmt.Printf("%#v\n", CoadingCourse)
	} else {
		fmt.Println("JSON was not Valid")
	}

	// some cases where you want to add data to key value pair
  // order is not guarented in the key value pair.
	
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}

/*
Concept:

- json.MarshalIndent() in Go is used to convert (marshal) a Go value into a nicely formatted JSON string, with indentation for readability.

Syntax:

json.MarshalIndent(v interface{}, prefix string, indent string)


| Parameter | Description                                                              |
| --------- | ------------------------------------------------------------------------ |
| `v`       | The Go value (struct, map, slice, etc.) you want to encode to JSON.      |
| `prefix`  | A string to prefix each line of JSON output (usually empty string `""`). |
| `indent`  | The string to use for indentation (e.g. `"\t"` or `"  "`).               |

It returns:

([]byte, error)


Comparison

- json.Marshal() → produces compact JSON (no spaces or indents).

- json.MarshalIndent() → produces pretty-printed, human-readable JSON.


Example:

package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Tags     []string `json:"tags"`
}

func main() {
	course := Course{
		Name:     "Golang Bootcamp",
		Price:    499,
		Platform: "LearnCodeOnline",
		Tags:     []string{"backend", "go", "web"},
	}

	// MarshalIndent converts struct to indented JSON
	finalJson, err := json.MarshalIndent(course, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

Output:


{
	"name": "Golang Bootcamp",
	"price": 499,
	"platform": "LearnCodeOnline",
	"tags": [
		"backend",
		"go",
		"web"
	]
}


*/