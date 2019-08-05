package main

import (
	//"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Creates the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	file := params["file"]

	fmt.Println(file)
}

// Renames the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {

}

// Deletes the file with the specified name
func deleteFile(w http.ResponseWriter, r *http.Request) {

}

// Establishes a tcp connection and sends a file
func getFile(w http.ResponseWriter, r *http.Request) {

}

// Modifies and stores the file according to the request
func updateFile(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/file/{file}", createFile).Methods("POST")
	r.HandleFunc("/api/file/{file}", createFile).Methods("PUT")
	r.HandleFunc("/api/file/{file}", deleteFile).Methods("DELETE")
	r.HandleFunc("/api/file/{file}", getFile).Methods("GET")
	r.HandleFunc("/api/file/{file}", updateFile).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8002", r))
}