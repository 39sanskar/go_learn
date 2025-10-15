package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Film model
type Film struct{
	ID        string     `json:"id"`
	Isbn      string     `json:"isbn"`  // ISBN stands for International Standard Book Number.
  Title     string     `json:"title"`
	Director  *Director  `json:"director"`
}

// Director model
type Director struct{
	Firstname string  `json:"firstname"`
  Lastname  string  `json:"lastname"`
}

var films []Film

// Get all films
func getAllFilms(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
	fmt.Println("Operation: GET all films") // <-- log
	json.NewEncoder(w).Encode(films)
}

// Get single film by ID
func getSingleFilm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Printf("Operation: GET film with ID=%s\n", params["id"])
	for _, item := range films {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Film not found", http.StatusNotFound)
}

// Create a new film
func createFilm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var film Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	film.ID = strconv.Itoa(rand.Intn(100000000)) // random ID and convert into string
	films = append(films, film) 
	fmt.Printf("Operation: POST create film '%s' with ID=%s\n", film.Title, film.ID)
	json.NewEncoder(w).Encode(film)

}

// Update an existing film
func updateFilm(w http.ResponseWriter, r *http.Request){
  // set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over the films, range
	// delete the film with the id that you i have sent
	// add a new film - the film that we send in the body of postman
  
	for index, item := range films {
		if item.ID == params["id"]{
			// Remove old film
			films = append(films[:index], films[index+1:]...)

			// Delete new data
			var updated Film 
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return 
			}

			// Keep the same ID
			updated.ID = params["id"]
			films = append(films, updated)
			fmt.Printf("Operation: PUT update film with ID=%s\n", params["id"])
			json.NewEncoder(w).Encode(updated)
			return 
		}
	}
	http.Error(w, "Film not found", http.StatusNotFound)
}

// Delete a Single film by ID
func deleteSingleFilm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range films {
		if item.ID == params["id"]{
			films = append(films[:index], films[index+1:]...)
			fmt.Printf("Operation: DELETE film with ID=%s\n", params["id"])
			json.NewEncoder(w).Encode(map[string]string{"message": "Film deleted successfully"})
			return 
		}
	}
	http.Error(w, "Film not found", http.StatusNotFound)
}

// Delete All films
func deleteAllFilms(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	// Reset the slice
	count := len(films)
	films = []Film{}
  fmt.Println("Operation: DELETE ALL films")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "All films deleted successfully",
		"deletedCount": count,
	})
}

func main(){
	// Initialize router
	r := mux.NewRouter()

	// Mock data
	films = append(films,
		Film{ID: "1", Isbn: "101", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
		Film{ID: "2", Isbn: "102", Title: "Game of Thrones", Director: &Director{Firstname: "David", Lastname: "Benioff"}},
		Film{ID: "3", Isbn: "103", Title: "Breaking Bad", Director: &Director{Firstname: "Vince", Lastname: "Gilligan"}},
		Film{ID: "4", Isbn: "104", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
		Film{ID: "5", Isbn: "105", Title: "Stranger Things", Director: &Director{Firstname: "The", Lastname: "Duffer Brothers"}},
		Film{ID: "6", Isbn: "106", Title: "Interstellar", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
		Film{ID: "7", Isbn: "107", Title: "Peaky Blinders", Director: &Director{Firstname: "Steven", Lastname: "Knight"}},
		Film{ID: "8", Isbn: "108", Title: "The Witcher", Director: &Director{Firstname: "Lauren", Lastname: "Hissrich"}},
		Film{ID: "9", Isbn: "109", Title: "The Lord of the Rings", Director: &Director{Firstname: "Peter", Lastname: "Jackson"}},
		Film{ID: "10", Isbn: "110", Title: "Money Heist", Director: &Director{Firstname: "Álex", Lastname: "Pina"}},
	)
	
  // Routes 
	r.HandleFunc("/films", getAllFilms).Methods(http.MethodGet)
	r.HandleFunc("/films/{id}", getSingleFilm).Methods(http.MethodGet)
	r.HandleFunc("/films", createFilm).Methods(http.MethodPost)
	r.HandleFunc("/films/{id}", updateFilm).Methods(http.MethodPut)
	r.HandleFunc("/films/{id}", deleteSingleFilm).Methods(http.MethodDelete)
	r.HandleFunc("/films", deleteAllFilms).Methods(http.MethodDelete)

	// Start server
	fmt.Println("🎬 Film API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

