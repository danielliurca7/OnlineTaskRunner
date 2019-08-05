package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

// Sends request to create the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	file := params["file"]

	response, err := http.Post("http://localhost:8002/api/file/" + file, "application/json", nil)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

// Sends request to rename the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {

}

// Sends request to delete the file with the specified name
func deleteFile(w http.ResponseWriter, r *http.Request) {

}

// Register if the client wrote some code
// Update the buffer and notify observers
func registerChange(w http.ResponseWriter, r *http.Request) {

}

// After the buffer of changes is full
// Send all of it to the io microservice
func commitChanges() {

}

// If the client wants to download the file
// The client microservice verify that he has permissions for the file
// And gives the client a port to listen to for the io microservice
func verifyConnection(w http.ResponseWriter, r *http.Request) {

}

// Make a request to the compute microservice
// To compile a workspace if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {

}

// Make a request to the compute microservice
// To run an executable/interpretable code
func runRequest(w http.ResponseWriter, r *http.Request) {

}

// Creates a custom test for a certain workspace
func createCustomTest(w http.ResponseWriter, r *http.Request) {

}

// Edits the build file
func editBuild(w http.ResponseWriter, r *http.Request) {

}

// Edits a task file
func editTasks(w http.ResponseWriter, r *http.Request) {

}

// Edits the docker image
func editDockerImage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/file/{file}", createFile).Methods("POST")
	r.HandleFunc("/api/file/{file}", renameFile).Methods("PUT")
	r.HandleFunc("/api/file/{file}", deleteFile).Methods("DELETE")
	r.HandleFunc("/api/file/{file}", verifyConnection).Methods("GET")

	r.HandleFunc("/api/change/{file}", registerChange).Methods("PUT")

	r.HandleFunc("/api/request", buildRequest).Methods("PUT")
	r.HandleFunc("/api/request", runRequest).Methods("GET")

	r.HandleFunc("/api/test/{name}", createCustomTest).Methods("POST")
	r.HandleFunc("/api/build", editBuild).Methods("PUT")
	r.HandleFunc("/api/tasks", editTasks).Methods("PUT")
	r.HandleFunc("/api/image", editDockerImage).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}