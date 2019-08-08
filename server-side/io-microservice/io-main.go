package main

import (
	"os"
	"encoding/json"
	"path/filepath"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"../utils"
	"../data_structures/file"
)

const fileSystem = "D:\\Projects\\OnlineTaskRunner\\server-side\\file_system\\"

// Creates the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var file file.File
	json.Unmarshal(body, &file)

	log.Println("Create file request for path " + filepath.Join(file.Path, file.Name))

	path := filepath.Join(fileSystem, file.Path)

	if !utils.FileExists(path) {
		err := os.MkdirAll(path, os.ModeDir)

		if err != nil {
			log.Println(err)
			utils.WriteResponse(w, 500, []byte("Folder could not be created"))
			return
		}
	}

	path = filepath.Join(path, file.Name)

	if utils.FileExists(path) {
		utils.WriteResponse(w, 400, []byte("File already exists at specified path"))
	} else {
		file, err := os.Create(path)

		defer file.Close()

		if err != nil {
			log.Println(err.Error())
			utils.WriteResponse(w, 500, []byte("File could not be created"))
			return
		}
	
		utils.WriteResponse(w, 200, []byte("File created"))
	}
}

// Renames the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var files []file.File
	json.Unmarshal(body, &files)

	oldFileName := filepath.Join(fileSystem, filepath.Join(files[0].Path, files[0].Name))
	newFileName := filepath.Join(fileSystem, filepath.Join(files[1].Path, files[1].Name))

	log.Println("Rename file request for path " + filepath.Join(files[0].Path, files[0].Name))
	
	if !utils.FileExists(oldFileName)  {
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
	body := utils.GetRequestBody(r)

	var file file.File
	json.Unmarshal(body, &file)

	log.Println("Delete file request for path " + filepath.Join(file.Path, file.Name))

	path := filepath.Join(fileSystem, filepath.Join(file.Path, file.Name))

	if !utils.FileExists(path) {
		utils.WriteResponse(w, 400, []byte("No file at the specified path"))
	} else {
		err := os.Remove(path)

		if err != nil {
			log.Println(err.Error())
			utils.WriteResponse(w, 500, []byte("File could not be deleted"))
			return
		}
	
		utils.WriteResponse(w, 200, []byte("File deleted"))
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

	r.HandleFunc("/api/file", createFile).Methods("POST")
	r.HandleFunc("/api/file", renameFile).Methods("PUT")
	r.HandleFunc("/api/file", deleteFile).Methods("DELETE")
	r.HandleFunc("/api/file", getFile).Methods("GET")
	r.HandleFunc("/api/file", updateFile).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8002", r))
}