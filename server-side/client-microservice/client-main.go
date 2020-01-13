package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../data_structures/credentials"
	"../utils"
)

func validateClaim() {

}

func authenticate(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var credentials credentials.Credentials
	json.Unmarshal(body, &credentials)

	log.Println("Authenticate request for user", credentials.Username)

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://db:10000/api/authenticate", "POST", body)
	utils.CheckError(err)

	data := utils.GetResponseBody(response)

	utils.WriteResponse(w, 200, data)
}

func getType(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var username struct{ Username string }
	json.Unmarshal(body, &username)

	log.Println("Get type request for user", username.Username)

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://db:10000/api/type", "POST", body)
	utils.CheckError(err)

	data := utils.GetResponseBody(response)

	utils.WriteResponse(w, 200, data)
}

// Make a request to the compute microservice
// To get the workspace files
func getFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Get file tree request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://cache:8000/api/get", "POST", body)

	utils.ForwardResponse(w, response, err)
}

// Sends request to create the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Create file tree request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://cache:8000/api/create", "POST", body)

	utils.ForwardResponse(w, response, err)
}

// Sends request to rename the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Rename file tree request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://cache:8000/api/rename", "POST", body)

	utils.ForwardResponse(w, response, err)
}

// Sends request to delete the file with the specified name
func deleteFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Delete file tree request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://cache:8000/api/delete", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func commitFiles(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Commit file tree request", string(body))

	// Make a request to the cache microservice
	response, err := utils.MakeRequest("http://cache:8000/api/commit", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func updateFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Update file tree request", string(body))

	// Make a request to the cache microservice
	response, err := utils.MakeRequest("http://cache:8000/api/update", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func getFileTree(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Get file tree request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://cache:8000/api/filetree", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func clearWorkspace(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Clear workspace request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://cache:8000/api/clear", "POST", body)

	utils.ForwardResponse(w, response, err)
}

// Make a request to the compute microservice
// To compile a workspace if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Build request for workspace ", string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://compute:9000/api/build", "POST", body)

	utils.ForwardResponse(w, response, err)
}

// Make a request to the compute microservice
// To run an executable/interpretable code
func runRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Run request for workspace", string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://compute:9000/api/run", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/authenticate", authenticate).Methods("POST")
	r.HandleFunc("/api/type", getType).Methods("POST")

	r.HandleFunc("/api/get", getFile).Methods("POST")
	r.HandleFunc("/api/create", createFile).Methods("POST")
	r.HandleFunc("/api/rename", renameFile).Methods("POST")
	r.HandleFunc("/api/delete", deleteFile).Methods("POST")
	r.HandleFunc("/api/commit", commitFiles).Methods("POST")
	r.HandleFunc("/api/update", updateFile).Methods("POST")
	r.HandleFunc("/api/filetree", getFileTree).Methods("POST")
	r.HandleFunc("/api/clear", clearWorkspace).Methods("POST")

	r.HandleFunc("/api/build", buildRequest).Methods("POST")
	r.HandleFunc("/api/run", runRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":7000", r))
}
