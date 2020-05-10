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
	r.HandleFunc("/api/student/{name:.*}", callbacks.GetStudentCourses).Methods("GET")
	r.HandleFunc("/api/assistant/{name:.*}", callbacks.GetAssistantCourses).Methods("GET")
	r.HandleFunc("/api/professor/{name:.*}", callbacks.GetProfessorCourses).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", r))
}
