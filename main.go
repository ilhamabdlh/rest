package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ilhamabdlh/rest/helper"
	"github.com/ilhamabdlh/rest/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func getDivisions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var divisions []models.DataOn

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var division models.DataOn
		// & character returns the memory address of the following variable.
		err := cur.Decode(&division) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		divisions = append(divisions, division)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(divisions) // encode similar to serialize process.
}

func getDivision(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var division models.DataOn
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	div, _ := primitive.ObjectIDFromHex(params["division"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"division": div}
	err := collection.FindOne(context.TODO(), filter).Decode(&division)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(division)
}

func createDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var division models.DataOn

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&division)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), division)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updateDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	div, _ := primitive.ObjectIDFromHex(params["id"])

	var division models.DataOn

	// Create filter
	//filter := bson.M{"_id": div}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&division)

	// prepare update model.
	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"division", bson.D{
	// 			{"firstname", division.Division.DivisionTwo},
	// 		}},
	// 	}},
	// }

	// err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&division)

	// if err != nil {
	// 	helper.GetError(err, w)
	// 	return
	// }

	division.ID = div

	json.NewEncoder(w).Encode(division)
}

func deleteDivision(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	div, err := primitive.ObjectIDFromHex(params["_id"])

	// prepare filter.
	filter := bson.M{"_id": div}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

// var client *mongo.Client


func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api", getDivisions).Methods("GET")
	r.HandleFunc("/api/{division}", getDivision).Methods("GET")
	r.HandleFunc("/api", createDivision).Methods("POST")
	r.HandleFunc("/api/{division}", updateDivision).Methods("PUT")
	r.HandleFunc("/api/{division}", deleteDivision).Methods("DELETE")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}