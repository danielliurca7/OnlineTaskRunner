package main

import (
	"log"
	"net/http"

	"./callbacks"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Files service is running")

	go callbacks.ListenForUpdates()

	r := mux.NewRouter()

	r.HandleFunc("/api/files", callbacks.GetFiles).Methods("GET")
	r.HandleFunc("/api/commit", callbacks.CommitFiles).Methods("POST")
	r.HandleFunc("/api/create", callbacks.CreateFile).Methods("POST")
	r.HandleFunc("/api/delete", callbacks.DeleteFile).Methods("DELETE")
	r.HandleFunc("/api/rename", callbacks.RenameFile).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", r))
}
