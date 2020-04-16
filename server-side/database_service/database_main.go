package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./callbacks"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/authenticate", callbacks.Authenticate).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", r))
}
