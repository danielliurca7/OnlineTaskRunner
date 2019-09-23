package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"../data_structures/containers"
	"../data_structures/file"
	"../data_structures/fileinfo"
	"../data_structures/token"
	"../data_structures/user"
	"../data_structures/workspace"
	"../utils"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	//"github.com/dgrijalva/jwt-go"
)

var tokens = make(map[string]string)
var redisClient *redis.Client

func authenticate(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var user user.User
	json.Unmarshal(body, &user)

	log.Println("Authenticate request for user", user.Username)

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:9000/api/authenticate", "POST", body)
	utils.CheckError(err)

	data := utils.GetResponseBody(response)

	var token token.Token
	json.Unmarshal(data, &token)

	tokens[token.Token] = token.PrivateToken

	utils.WriteResponse(w, 200, data)
}

// Sends request to create the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	fi := utils.GetFileinfoFromBody(body)

	log.Println("Create file request for", utils.GetPath(fi))

	key := utils.GetWorkspaceHash(fi.Workspace)

	val, err := redisClient.Get(key).Result()
	utils.CheckError(err)

	var files []file.File
	json.Unmarshal([]byte(val), &files)

	files = append(files, file.File{
		Path:  filepath.Join(fi.Path...),
		Data:  "",
		IsDir: false,
	})

	b, err := json.Marshal(files)
	utils.CheckError(err)

	err = redisClient.Set(key, b, 0).Err()
	utils.CheckError(err)

	utils.WriteResponse(w, 200, []byte("File created"))
}

// Sends request to create the folder with the specified name
func createFolder(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	fi := utils.GetFileinfoFromBody(body)

	log.Println("Create file request for", utils.GetPath(fi))

	key := utils.GetWorkspaceHash(fi.Workspace)

	val, err := redisClient.Get(key).Result()
	utils.CheckError(err)

	var files []file.File
	json.Unmarshal([]byte(val), &files)

	files = append(files, file.File{
		Path:  filepath.Join(fi.Path...),
		Data:  "",
		IsDir: true,
	})

	b, err := json.Marshal(files)
	utils.CheckError(err)

	err = redisClient.Set(key, b, 0).Err()
	utils.CheckError(err)

	utils.WriteResponse(w, 200, []byte("Folder created"))
}

// Sends request to rename the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var fis []fileinfo.Fileinfo
	json.Unmarshal(body, &fis)

	log.Println("Rename file request from", utils.GetPath(fis[0]),
		"to", utils.GetPath(fis[1]))

	key := utils.GetWorkspaceHash(fis[0].Workspace)

	val, err := redisClient.Get(key).Result()
	utils.CheckError(err)

	var files []file.File
	json.Unmarshal([]byte(val), &files)

	path := filepath.Join(fis[0].Path...)

	for i, v := range files {
		if path == v.Path {
			files[i].Path = filepath.Join(fis[1].Path...)
			break
		}
	}

	b, err := json.Marshal(files)
	utils.CheckError(err)

	err = redisClient.Set(key, b, 0).Err()
	utils.CheckError(err)

	utils.WriteResponse(w, 200, []byte("File renamed"))
}

// Sends request to delete the file with the specified name
func deleteFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Delete file request", utils.GetPath(utils.GetFileinfoFromBody(body)))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:10000/api/delete", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func registerChange(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Get file tree request", string(body))

	// Make a request to the cache microservice
	response, err := utils.MakeRequest("http://localhost:7000/api/filetree", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func getFileTree(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Get file tree request", string(body))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:10000/api/filetree", "POST", body)

	utils.ForwardResponse(w, response, err)
}

// If the client wants to download the file
// The client microservice verify that he has permissions for the file
// And gives the client a port to listen to for the io microservice
func verifyConnection(w http.ResponseWriter, r *http.Request) {

}

// Make a request to the compute microservice
// To get the workspace files
func getRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var ws workspace.Workspace
	json.Unmarshal(body, &ws)

	log.Println("Get request for workspace ", string(body))

	key := utils.GetWorkspaceHash(ws)

	_, err := redisClient.Get(key).Result()

	if err != nil {
		response, err := utils.MakeRequest("http://localhost:10000/api/workspace", "POST", body)

		b := utils.GetResponseBody(response)

		err = redisClient.Set(key, b, 0).Err()
		utils.CheckError(err)
	}
}

// Make a request to the compute microservice
// To compile a workspace if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Build request for workspace ", string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://localhost:8001/api/build", "PUT", body)

	utils.ForwardResponse(w, response, err)
}

// Make a request to the compute microservice
// To run an executable/interpretable code
func runRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Run request for workspace", string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://localhost:8001/api/run", "GET", body)

	utils.ForwardResponse(w, response, err)
}

// Make a request to the compute microservice
// To clean the workspace
func cleanRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Clean request for workspace", string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://localhost:8001/api/clean", "DELETE", body)

	utils.ForwardResponse(w, response, err)
}

// Make a request to the compute microservice
// To delete the workspace
func clearRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var ws workspace.Workspace
	json.Unmarshal(body, &ws)

	log.Println("Clear request for", ws)

	key := utils.GetWorkspaceHash(ws)

	val, err := redisClient.Get(key).Result()

	var c containers.WorkspaceContainer

	c.Workspace = ws
	json.Unmarshal([]byte(val), &c.Files)

	b, err := json.Marshal(c)

	response, err := utils.MakeRequest("http://localhost:10000/api/update", "POST", b)

	err = redisClient.Del(key).Err()
	utils.CheckError(err)

	utils.ForwardResponse(w, response, err)
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	/*err := redisClient.Set("key", "value1", 0).Err()
	utils.CheckError(err)

	val, err := redisClient.Get("key").Result()
	utils.CheckError(err)
	log.Println("key", val)

	err = redisClient.Set("key", "value2", 0).Err()
	utils.CheckError(err)

	val, err = redisClient.Get("key").Result()
	utils.CheckError(err)
	log.Println("key", val)

	err = redisClient.Del("key").Err()
	utils.CheckError(err)

	val, err = redisClient.Get("key").Result()
	utils.CheckError(err)
	log.Println("key", val)*/

	r := mux.NewRouter()

	r.HandleFunc("/api/authenticate", authenticate).Methods("POST")

	r.HandleFunc("/api/create_file", createFile).Methods("POST")
	r.HandleFunc("/api/create_folder", createFolder).Methods("POST")
	r.HandleFunc("/api/rename", renameFile).Methods("POST")
	r.HandleFunc("/api/delete", deleteFile).Methods("POST")
	r.HandleFunc("/api/file", verifyConnection).Methods("GET")

	r.HandleFunc("/api/change", registerChange).Methods("POST")
	r.HandleFunc("/api/filetree", getFileTree).Methods("POST")

	r.HandleFunc("/api/get", getRequest).Methods("POST")
	r.HandleFunc("/api/build", buildRequest).Methods("POST")
	r.HandleFunc("/api/run", runRequest).Methods("POST")
	r.HandleFunc("/api/clean", cleanRequest).Methods("POST")
	r.HandleFunc("/api/clear", clearRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":6000", r))
}
