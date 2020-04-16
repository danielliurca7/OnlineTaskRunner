package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./callbacks"
)

func main() {
	log.Println("Compute service is running")

	r := mux.NewRouter()

	r.HandleFunc("/api/build", callbacks.BuildImage).Methods("POST")
	r.HandleFunc("/api/run", callbacks.RunContainer).Methods("POST")
	r.HandleFunc("/api/stop", callbacks.StopContainer).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", r))
}
