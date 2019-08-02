package compute-microservice

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Runs the project in a workspace
// Compiles if necessary
// A run script will be included
func getRunResult() {

}

// Compile files if necessary
func compileFiles() {

}

// Run a specific test
func runTest() {

}

// Get the file to run from specified microservice
func getFile() {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/run/{workspace}", getRunResult).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}