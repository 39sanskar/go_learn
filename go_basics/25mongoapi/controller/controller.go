package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/39sanskar/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://coadinghub<password>@cluster1.opx0mjp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1"

// real <password> => sanskar 

const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
  fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file 

// insert 1 record 

// Note => here i have use insertOneMovie in i has a lowercase because it is expected this is helper method i will never be exporting this method, this is the mongodb helpher method the other method that we talked about they will be written with the first upper letter because their are chances it is 100 percent sure we will be exporting those method that other people also can use that.
func insertOneMovie(movie model.Netflix){
  inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
    log.Fatal(err)
	}
  fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	// everything inside the mongodb is not json technically it is bson. but bson provides you more thing like objectId.
  filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
    log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount 
  
}

// get all movies from database
func getAllMovies() []primitive.M{
	// here cur => cursor {mongodb cursor}
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M 
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
  return movies
}

// Actual controller - file 
// in the router file i need this controller.
// r => reader gives the reference of all the Request that coming in. 
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies) // json is come from the frontend.
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
  insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
  
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
