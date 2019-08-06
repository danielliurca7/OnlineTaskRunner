package main

import (
	"os"
	//"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"../utils"
)

const fileSystem = "D:\\Projects\\OnlineTaskRunner\\server-side\\file_system\\"

func fileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

// Creates the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fileName := params["file"]

	path := fileSystem + fileName

	log.Println("Create file request for path " + path)

	if !fileExists(path) {
		file, err := os.Create(path)

		if err != nil {
			fmt.Println(err.Error())

			utils.WriteResponse(w, 500, []byte("File could not be created"))
		}

		defer file.Close()

		utils.WriteResponse(w, 200, []byte("File created"))
	} else {
		utils.WriteResponse(w, 400, []byte("File already exists at specified path"))
	}	
}

// Renames the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	oldFileName := fileSystem + params["file"]

	body := utils.GetRequestBody(r)
	newFileName := fileSystem + string(body)

	log.Println("Rename file request for path " + oldFileName)
	
	if !fileExists(oldFileName)  {
			utils.WriteResponse(w, 404, []byte("File to rename not found"))
	} else {
		err := os.Rename(oldFileName, newFileName)

		if err != nil {
			log.Println(err.Error())

			utils.WriteResponse(w, 500, []byte("File could not be renamed"))
		} else {
			utils.WriteResponse(w, 200, []byte("File renamed"))
		}
	}
}

// Deletes the file with the specified name
func deleteFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fileName := params["file"]

	path := fileSystem + fileName

	log.Println("Delete file request for path " + path)

	if !fileExists(path) {
		utils.WriteResponse(w, 404, []byte("File to delete not found"))
	} else {
		err := os.Remove(path)

		if err != nil {
			fmt.Println(err.Error())

			utils.WriteResponse(w, 500, []byte("File could not be deleted"))
		} else {
			utils.WriteResponse(w, 200, []byte("File deleted"))
		}
	}
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
	r.HandleFunc("/api/file/{file}", renameFile).Methods("PUT")
	r.HandleFunc("/api/file/{file}", deleteFile).Methods("DELETE")
	r.HandleFunc("/api/file/{file}", getFile).Methods("GET")
	r.HandleFunc("/api/file/{file}", updateFile).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8002", r))
}