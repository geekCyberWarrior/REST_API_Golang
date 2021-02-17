package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Email    string  `json:"email"`
}

var students []Student

func db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")// Connect to //MongoDB
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	fmt.Println("Connected to MongoDB!")
	return client
}

var studentCollection = db().Database("DB_NAME").Collection("students")

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)

	insertResult, _ := studentCollection.InsertOne(context.TODO(),student)
	json.NewEncoder(w).Encode(insertResult.InsertedID)
	
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, _ := studentCollection.Find(context.TODO(), bson.M{})
	var students []bson.M
	cursor.All(context.TODO(), &students)

	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var student primitive.M
	studentCollection.FindOne(context.TODO(), bson.D{{"name", params["name"]}}).Decode(&student)
	
	json.NewEncoder(w).Encode(student)

}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	studentCollection.ReplaceOne(
	    context.TODO(),
	    bson.M{"name": params["name"]},
	    bson.M{
	        "name":  student.Name,
	        "email": student.Email,
	        "phone": student.Phone})

	json.NewEncoder(w).Encode(student)

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	opts := options.Delete().SetCollation(&options.Collation{})
	res, _ := studentCollection.DeleteOne(context.TODO(), bson.M{"name": params["name"]}, opts)
	json.NewEncoder(w).Encode(res.DeletedCount)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{name}", getStudent).Methods("GET")
	r.HandleFunc("/students/{name}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{name}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", r))	
}
