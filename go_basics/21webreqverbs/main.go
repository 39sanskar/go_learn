package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

// Make sure server is up and running (gowebserver)
func PerformGetRequest() {
	// Go’s visibility rule for identifiers.
	// The first letter of PerformGetRequest being uppercase means the function is exported (public).Uppercase first letter → exported (public) — visible outside the package.
	// If you write it as performGetRequest, it becomes unexported (private).Lowercase first letter → unexported (private) — visible only within the same package.

	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder 
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String()) // whaterer data holding inside it will convert into a format of string.

  // fmt.Println(content) // this will print is a byte format
	// fmt.Println(string(content))

}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	// fake json payload (payload is another name if you send some data.)

	requestBody := strings.NewReader(`
	    {
	      "coursename": " Let's go with golang ",
				"price": 0,
				"platform": " CoadingHub.in "
	    }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
  defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest(){
	 const myurl = "http://localhost:8000/postform"

	 // formdata

	 data := url.Values{}
	 data.Add("firstname", "sanskar")
	 data.Add("lastname", "mishra")
	 data.Add("email", "sankar@gmail.com")

	 response, err := http.PostForm(myurl, data)
	 if err != nil {
		  panic(err)
	 }

	 defer response.Body.Close() // when everything is done we are going to defer a request for close the request.

	 content, _ := ioutil.ReadAll(response.Body)
	 fmt.Println(string(content))
}