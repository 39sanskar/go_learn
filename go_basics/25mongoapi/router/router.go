package router

import (
	"github.com/39sanskar/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
  router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
  router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
  
	return router 
}

/*

-	func Router() *mux.Router {

	}

Explanation:
- mux.Router is a struct type provided by the Gorilla Mux package.
- The * (asterisk) before it means your function returns a pointer to that struct — i.e., *mux.Router means “pointer to a Router”.

So:
- mux.Router → the actual type (a struct)
- *mux.Router → pointer to that type

- So router := mux.NewRouter() gives you a *mux.Router.
- Returning a pointer:
- avoids copying the entire router struct,
- ensures any middleware or route you add later affects the same instance.

*/