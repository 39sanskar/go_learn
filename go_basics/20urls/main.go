package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing the url
	result, _  := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())  // Port it is method not a property so using Port().
  // fmt.Println(result.RawQuery)

	qparams := result.Query() // stores all these parameter into better format.
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// in this case remember this is very essential things you always pass on the reference  & in this case. 
	partsOfUrl := &url.URL{
		  Scheme: "https",
      Host: "lco.dev",
			Path: "/tutcss",
			RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
