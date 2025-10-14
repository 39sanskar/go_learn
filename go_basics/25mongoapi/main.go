package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/39sanskar/mongoapi/router"
)

func main(){
  fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
  fmt.Println("Listening on port 4000...")

}



// Note => in the application which are design in golang it is aspected in the root directory of the project just one go file it can be main.go file it can be other go file but it should be just one main.go file. if i am creating more go file in the root directory like controllers.go and all of that, they are going to create an issue so i need to create a folder and inside those folder i have to create more go files.

/*

// for database connection refer  go.mongodb.org 

- Important Command

1. first initialise the  go mod init github.com/39sanskar/mongoapi => 
What It Does
- Initializes a new Go module in your project.
	- It creates a file called go.mod in your project root.
	- This file tells Go that your project is a module and tracks dependencies.
- Sets the module path
	- github.com/39sanskar/mongoapi becomes the module path.
	- This path is used when importing your own packages in the project or when other projects import your module.
- Enables dependency management
	- After initializing, you can use go get to add external packages.
	- Go will track them in go.mod and go.sum.

2. go get -u github.com/gorilla/mux 

3. go build main.go 
*/
