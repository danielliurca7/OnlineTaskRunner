package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"../data_structures/containers"
	"../data_structures/file"
	"../utils"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var redisClient *redis.Client

func deleteFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Delete file request", utils.GetPath(utils.GetFileinfoFromBody(body)))

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:10000/api/delete", "POST", body)

	utils.ForwardResponse(w, response, err)
}

func registerChange(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.OneChangeContainer
	json.Unmarshal(body, &c)

	log.Println("Register change request for", c.Fileinfo)

	key := utils.GetWorkspaceHash(c.Fileinfo.Workspace)

	val, _ := redisClient.Get(key).Result()

	b := []byte(val)

	filename := filepath.Join(c.Fileinfo.Path...)

	var files []file.File
	json.Unmarshal(b, &files)

	for i, v := range files {
		if filename == v.Path {
			log.Println(files[i].Data)

			files[i].Data = string(utils.ApplyChange([]byte(v.Data), c.Change))

			log.Println(files[i].Data)

			b, err := json.Marshal(files)

			err = redisClient.Set(key, b, 0).Err()
			utils.CheckError(err)
		}
	}
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

	r.HandleFunc("/api/create_file", createFile).Methods("POST")
	r.HandleFunc("/api/create_folder", createFolder).Methods("POST")
	r.HandleFunc("/api/rename", renameFile).Methods("POST")
	r.HandleFunc("/api/delete", deleteFile).Methods("POST")
	r.HandleFunc("/api/get_files", getFiles).Methods("POST")
	r.HandleFunc("/api/change", registerChange).Methods("POST")

	log.Fatal(http.ListenAndServe(":7000", r))
}
