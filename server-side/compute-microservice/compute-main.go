package compute_microservice

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Compile files if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {

}

// Runs the project in a workspace
// Compiles if necessary
// A build script will be included
func runRequest(w http.ResponseWriter, r *http.Request) {

}

// Run a specific test
func registerChange(w http.ResponseWriter, r *http.Request) {

}

// Get the file to run from specified microservice
func getFile() {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/build", getRunResult).Methods("PUT")
	r.HandleFunc("/api/run", getRunResult).Methods("GET")
	r.HandleFunc("/api/change/{file}", registerChange).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8001", r))
}