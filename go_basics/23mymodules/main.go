package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()   // this  mux.NewRouter is actually handle the HandleFunction 
	// this mux.NewRouter() automatically bringing this "github.com/gorilla/mux"

	r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":4000", r) // this actual code throw some error in the case of web we actually do the classic comma ok syntax.

	log.Fatal(http.ListenAndServe(":4000", r)) // log.Fatal() which automatically something goes fail it just log that's value.

}

func greeter() {
	fmt.Println("Hey there mod users")
}

// see when you write backend code.
func serveHome(w http.ResponseWriter, r *http.Request){
	  // if you want to send some response that is done through w.Write 
	  w.Write([]byte("<h1>Welcome to golang series on Coading hub</h1>"))
}

/*

- Let's start registering a couple of URL paths and handlers:

func main(){
  r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r) 
}


- Remember => whenever we use go mod these are expensive operations,
- when you run  go mod tidy  so indirect line are gone from the go.mod file.

- Important Command

go get -u github.com/gorilla/mux  => it going to go ahead and downloading this one here.

go env => it show all the variables.
go build . => it build everything inside this one.
go mod verify => this is going to say all modules are verified.

go list => Prints the import path of the current module’s package (e.g. github.com/username/project).

go list all => lists all packages that are:
in your module
in its dependencies
and in the Go standard library


go list -m all => ist's all modules that your project depends on —
including:
your own module,
direct dependencies (what you imported yourself), and
indirect dependencies (dependencies of your dependencies).


go list -m -versions github.com/gorilla/mux => This command lists all the available tagged versions of a Go module from its source repository (for example, GitHub).

go mod tidy => tidys up all the library that you have depending on that it also removes all those packages that you are not using it also try to bring in all packages maybe some region remove that.

go mod why github.com/gorilla/mux  => it shows this is the module  (github.com/39sanskar/mymodules) is depending on this module (github.com/gorilla/mux).

go mod graph  => Shows who depends on whom.

go mod edit -go 1.24.3 => actually edit the go version

go mod vendor => vendor folder is consider as node_modules, now it is not bringing up everything from web it is going to just bring it at one time you can package this entirly  and everything now it is going to run through this vendor not directly we have to pass on a flag.

- if you do this  go run main.go  obsiously this is going to bring up everything from the web itself. 

- but if you do   go run -mod=vendor main.go   then its going to first look into vendor folder  if their is one it's going to bring up everything from here and i even haven't yet scratched  the surface of what go mod actually has to bring in so you can see there is a versioning.


- go.sum 
- What is go.sum?
- When you use Go modules (i.e. your project has a go.mod file), Go tracks the dependencies your code needs.

- The go.mod file records which dependencies (and versions) your project uses.
- The go.sum file records checksums (cryptographic hashes) of those dependencies to make sure they haven't been tampered with.

- go.sum = list of verified dependency fingerprints
- go.mod = list of required module versions

- Why is it Important?
- Ensures security — detects if a dependency was modified or replaced maliciously.
- Ensures reproducibility — every team member builds with the exact same dependency versions.
- Used automatically by Go commands like go build, go test, etc.

- Common Commands Related to It

| Command              | Purpose                                                        |
| -------------------- | -------------------------------------------------------------- |
| `go mod tidy`        | Cleans unused deps & updates `go.sum`                          |
| `go mod verify`      | Checks that all downloaded modules match their `go.sum` hashes |
| `go clean -modcache` | Clears cached modules if something seems broken                |
| `go list -m all`     | Lists all modules used                                         |
| `go mod why module`  | Shows  why  a module is included                               |
| `go mod graph`       | Shows  who depends on whom*                                    |


*/
