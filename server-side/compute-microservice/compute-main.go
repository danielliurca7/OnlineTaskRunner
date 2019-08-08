package main

import (
	//"os"
	"os/exec"
	"bytes"
	//"encoding/json"
	"path/filepath"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/rs/xid"
	"../utils"
)

const tmp = "D:\\Projects\\OnlineTaskRunner\\server-side\\tmp"

/*
Logic for hashing the directory names

var m = make(map[[64]byte]string)

TO BE CONTINUED

	rawBody := utils.GetRequestBody(r)
	var body [64]byte
	copy(body[:], rawBody[0:64])

	folder, ok := m[body]

	if !ok {
		folder = xid.New().String()
		path := filepath.Join(tmp, folder)

		err := os.Mkdir(path, os.ModeDir)

		if err != nil {
			log.Println(err)
			utils.WriteResponse(w, 500, []byte("Folder could not be created"))
			return
		}

		m[body] = folder

		utils.WriteResponse(w, 200, []byte("Folder created"))

		// Get files from io_MS
	}

	log.Println("Build request for folder " + folder)

	path := filepath.Join(tmp, folder)
	log.Println(path)
	*/


// Compile files if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {


	workspace := "test"
	path := filepath.Join(tmp, workspace)

	// Format the make build command
	cmd := exec.Command("make", "build")
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = path

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Build failed with", err)
		utils.WriteResponse(w, 400, e.Bytes())
	} else {
		utils.WriteResponse(w, 200, e.Bytes())
	}
}

// Runs the project in a workspace
func runRequest(w http.ResponseWriter, r *http.Request) {
	workspace := "test"
	path := filepath.Join(tmp, workspace)

	// Format the make run command
	cmd := exec.Command("make", "run")
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = path

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Run failed with", err)
		utils.WriteResponse(w, 400, e.Bytes())
	} else {
		utils.WriteResponse(w, 200, o.Bytes())
	}
}

// Cleans the project in a workspace
func cleanRequest(w http.ResponseWriter, r *http.Request) {
	workspace := "test"
	path := filepath.Join(tmp, workspace)

	// Format the make clean command
	cmd := exec.Command("make", "clean")
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = path

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Run failed with", err)
		utils.WriteResponse(w, 400, e.Bytes())
	} else {
		utils.WriteResponse(w, 200, o.Bytes())
	}
}

// Run a specific test
func registerChange(w http.ResponseWriter, r *http.Request) {

}

// Get the file to run from specified microservice
func getFile() {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/build", buildRequest).Methods("PUT")
	r.HandleFunc("/api/run", runRequest).Methods("GET")
	r.HandleFunc("/api/clean", cleanRequest).Methods("DELETE")
	r.HandleFunc("/api/change/{file}", registerChange).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8001", r))
}