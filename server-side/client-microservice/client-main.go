package client-microservice

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Creates the file with the specified name
func createFile() {

}

// Deletes the file with the specified name
func deleteFile() {

}

// Register if the client wrote some code
// Update the buffer and notify observers
func registerChange() {

}

// After the buffer of changes is full
// Send all of it to the io microservice
func commitChanges() {

}

// If the client wants to download the file
// The client microservice verify that he has permissions for the file
// And gives the client a port to listen to for the io microservice
func verifyConnection() {

}

// Make a request to the compute microservice
// To run an executable/interpretable code
func makeRunRequest() {

}

// Get's the result from the compute microservice
// And passes it to the user
func getResult() {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/result/{workspace}", getResult).Methods("GET")
	r.HandleFunc("/api/file/{filePath}", verifyConnection).Method("GET")
	r.HandleFunc("/api/file/{filePath}", registerChange).Methods("PUT")
	r.HandleFunc("/api/file/{filePath}", createFile).Methods("POST")
	r.HandleFunc("/api/file/{filePath}", deleteFile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}