package io-microservice

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Establishes a tcp connection and sends a file
func getFile() {

}

// Modifies and stores the file according to the request
func storeFile() {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/file/{filePath}", getFile).Methods("GET")
	r.HandleFunc("/api/file/{filePath}", storeFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", r))
}