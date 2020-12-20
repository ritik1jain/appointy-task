package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ritik1jain/go-rest-api/helper"
	"github.com/ritik1jain/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var db = helper.ConnectDB()

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var users []models.User

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := db.Collection("user").Find(context.TODO(), bson.M{})

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
		var user models.User
		// & character returns the memory address of the following variable.
		err := cur.Decode(&user) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users) // encode similar to serialize process.
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := db.Collection("user").FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Timestamp = time.Now()
	// insert our user model.
	result, err := db.Collection("user").InsertOne(context.TODO(), user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&contact)

	contact.Timestamp = time.Now()

	// insert our contact model.
	result, err := db.Collection("contact").InsertOne(context.TODO(), contact)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func listContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id,timestamp from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])
	// timestamp, _ := params["timestamp"]

	// var contact models.Contact

	// Create filter
	// filter := bson.M{"_id": id}
	// filter1 := bson.M{"_id2": id}

	// we created user array
	var users []models.User

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := db.Collection("contact").Find(context.TODO(), bson.M{"userid1": id})
	curr, eror := db.Collection("contact").Find(context.TODO(), bson.M{"userid2": id})

	if err != nil {
		if eror == nil {
			cur = curr
		} else {
			helper.GetError(eror, w)
			return
		}

	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var user models.User
		// & character returns the memory address of the following variable.
		err := cur.Decode(&user) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users) // encode similar to serialize process.

	// json.NewEncoder(w).Encode(contact)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/", listContacts).Methods("GET")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
