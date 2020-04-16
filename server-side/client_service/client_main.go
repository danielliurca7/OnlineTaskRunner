package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	compute "./callbacks/compute"
	database "./callbacks/database"
	files "./callbacks/files"
	"./utils"
)

func main() {
	log.Println("Client service is running")

	go files.ListenForUpdates()

	r := mux.NewRouter()

	r.Handle("/socket/change", utils.NewChangeServer())

	r.HandleFunc("/api/authenticate", database.Authenticate).Methods("GET")

	r.HandleFunc("/api/files", files.GetFiles).Methods("GET")
	r.HandleFunc("/api/commit", files.CommitFiles).Methods("POST")
	r.HandleFunc("/api/create", files.CreateFile).Methods("POST")
	r.HandleFunc("/api/delete", files.DeleteFile).Methods("DELETE")
	r.HandleFunc("/api/rename", files.RenameFile).Methods("PUT")
	r.HandleFunc("/api/update", files.UpdateFile).Methods("PUT")

	r.HandleFunc("/api/build", compute.BuildImage).Methods("POST")
	r.HandleFunc("/api/run", compute.RunContainer).Methods("POST")
	r.HandleFunc("/api/stop", compute.StopContainer).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}
