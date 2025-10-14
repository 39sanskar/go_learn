package main

import (
	"encoding/json"
	"math/rand"         // Required for rand.Seed and rand.Intn
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Creating Model for course - file (whenever we say -file this means suppose to go in a file but we are not putting it file)
type Course struct {
  CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author  *Author   `json:"author"` // this is where not i am referencing the variable directly i am just defining the types here (it is pointer type  *Author).
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func  (c *Course) IsEmpty() bool { // passing the structure of the course
	// return c.CourseId == "" && c.CourseName == ""

	// i don't want to check for courseId i want that user should be allow to move further if the courseId is not empty they are passing on this value. Why i am doing so because i want to manually create the courseId i don't want to relay on the user you are just going go ahead and pass me unique id.

	return c.CourseName == "" 
}
func main(){
  fmt.Println("API - CoadingHub.in")
	r := mux.NewRouter()

	// seeding 
  courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Sanskar Mishra", Website: "coding.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Sanskar Mishra", Website: "go.dev"}})

	// routing
  r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")
	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route 

// serveHome(w http.ResponseWriter, r *http.Request) => it is always been covered by two thing writer and reader
// reader is where you get the value from who ever is sending the Request.
// writer is where you write the response for that.

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by Coading Hub</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	// in some times in some cases you might want to set a special header a things up here responses definitely one of the things but some time you want to set extreme headers.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
  id := params["id"]
	// loop through courses, find matching id and return the response
	for _, course := range courses{
		if course.CourseId == id {
      json .NewEncoder(w).Encode(course)
			return 
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// 1.what if: body is empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send some data"})
		return 
	}

	// what about data which is send up like {}

	// 2.Decode request body into struct
	var course Course 
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
		return 
	}

	// 3.Check if course is empty
	if course.IsEmpty(){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "No data inside JSON"})
		return 
	}
  
	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON
  
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "Course title already exists"})
			return
		}
	}

	// 4.Assign ID and append to global slice
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// Append to slice
  courses = append(courses, course)

	// 5.Return created course
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
	
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req   
	params := mux.Vars(r)
  // somebody needs to give me the id that what is that course that i should update and also from the body itself i will be able to grab all the value in the json format.
	// this time two value are coming in one is from the request parameters from the url itself and another one inside the body.

	id := params["id"]
	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == id {
      // Remove the old course
			courses = append(courses[:index], courses[index+1:]...)

			// Decode the new course
			var updatedCourse Course
			if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
				return 
			}

			// Check for empty data
			if updatedCourse.IsEmpty(){
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Empty or missing fields in JSON"})
				return
			}

			// Keep same ID
			updatedCourse.CourseId = id 

			// Append back the new course
			courses = append(courses, updatedCourse)

			// Send updated response
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCourse)
			return
	  }
	}	
  // Todo: send a response when id is not found
	// If no match found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": fmt.Sprintf("Course with ID %s not found", id),
	})
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
  fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	deleted := false // flag to track if we deleted a course
	// Action plan => loop, match the id, remove (index, index+1)

	// Loop and delete by ID
	for index, course := range courses {
		if course.CourseId == id { // compare correctly
			// remove the course
			courses = append(courses[:index], courses[index+1:]...)
      deleted = true
			break // exit the loop
		}
	}
  
	// TODO: send a confirm or deny response

	// Send confirmation response 
	if deleted {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Course with ID %s deleted successfully", id),
		})
	} else {
		// if no course found
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Course with ID %s not found", id),
		})
	}
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request){
  fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")

	// Clear the slice
	courses = []Course{} // resets the slice to an empty slice.

	// Send confirmation response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "All courses have been deleted successfully",
	})
}

