package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	"../data_structures/containers"
	"../data_structures/file"
	"../data_structures/fileinfo"
	"../data_structures/tree"
	"../data_structures/workspace"
	"../utils"
)

var redisClient *redis.Client

func getFilesFromDisk(bytes []byte) error {
	f, err := redisClient.HGetAll(string(bytes)).Result()

	if len(f) != 0 {
		return errors.New("Files already in cache")
	}

	response, err := utils.MakeRequest("http://localhost:10000/api/workspace", "POST", bytes)
	if err != nil {
		return err
	}

	var files []file.File
	json.Unmarshal(utils.GetResponseBody(response), &files)

	for _, v := range files {
		fi := fileinfo.Fileinfo{
			Path:  v.Path,
			IsDir: v.IsDir,
		}

		b, err := json.Marshal(fi)
		if err != nil {
			return err
		}

		err = redisClient.HSet(string(bytes), string(b), v.Data).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func getFilesFromCache(bytes []byte) ([]byte, error) {
	var ws workspace.Workspace
	json.Unmarshal(bytes, &ws)

	c := containers.WorkspaceContainer{
		Workspace: ws,
	}

	files, err := redisClient.HGetAll(string(bytes)).Result()
	utils.CheckError(err)

	if len(files) == 0 {
		return nil, errors.New("Not found in cache")
	}

	for k, v := range files {
		var fi fileinfo.Fileinfo
		json.Unmarshal([]byte(k), &fi)

		file := file.File{
			Path:  fi.Path,
			Data:  v,
			IsDir: fi.IsDir,
		}

		c.Files = append(c.Files, file)
	}

	return json.Marshal(c)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Get request for workspace", string(body))

	var c containers.OneFileContainer
	json.Unmarshal(body, &c)

	b1, err := json.Marshal(c.Workspace)
	utils.CheckError(err)

	b2, err := json.Marshal(c.Fileinfo)
	utils.CheckError(err)

	log.Println(string(b1), string(b2))

	ok, err := redisClient.HExists(string(b1), string(b2)).Result()
	utils.CheckError(err)

	if !ok {
		err := getFilesFromDisk(b1)
		utils.CheckError(err)
	}

	data, err := redisClient.HGet(string(b1), string(b2)).Result()
	utils.CheckError(err)

	utils.WriteResponse(w, 200, []byte(data))
}

func getWorkspace(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Get request for workspace", string(body))

	err := getFilesFromDisk(body)
	utils.CheckError(err)

	bytes, err := getFilesFromCache(body)
	utils.CheckError(err)

	utils.WriteResponse(w, 200, bytes)
}

func getFiletree(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	fileNames, err := redisClient.HKeys(string(body)).Result()
	utils.CheckError(err)

	root := &tree.Node{}

	for _, file := range fileNames {
		presentNode := root

		var fi fileinfo.Fileinfo
		json.Unmarshal([]byte(file), &fi)

		for index, fileName := range fi.Path {
			found := false

			for _, node := range presentNode.Children {
				if node.Name == fileName {
					presentNode = node
					found = true
					break
				}
			}

			if !found {
				node := &tree.Node{
					Name:     fileName,
					Children: []*tree.Node{},
					IsDir:    fi.IsDir,
					Path:     fi.Path[:index+1],
				}
				presentNode.Children = append(presentNode.Children, node)
				presentNode = node
			}
		}
	}

	if err != nil {
		log.Println(err)
		utils.WriteResponse(w, 500, []byte("Could not walk through file tree"))
	}

	b, err := json.Marshal(root.Children)
	utils.CheckError(err)

	utils.WriteResponse(w, 200, b)
}

func clearWorkspace(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Clear request for workspace", string(body))

	err := redisClient.Del(string(body)).Err()
	utils.CheckError(err)
}

func createFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.OneFileContainer
	json.Unmarshal(body, &c)

	log.Println("Create request for file", c.Fileinfo, "in workspace", c.Workspace)

	b1, err := json.Marshal(c.Workspace)
	utils.CheckError(err)

	ok, err := redisClient.Exists(string(b1)).Result()
	utils.CheckError(err)

	if ok == 0 {
		err = getFilesFromDisk(b1)
		utils.CheckError(err)
	}

	b2, err := json.Marshal(c.Fileinfo)
	utils.CheckError(err)

	exists, err := redisClient.HExists(string(b1), string(b2)).Result()
	utils.CheckError(err)

	if !exists {
		err = redisClient.HSet(string(b1), string(b2), c.Data).Err()
		utils.CheckError(err)
	}
}

func renameFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.TwoFileinfoContainer
	json.Unmarshal(body, &c)

	log.Println("Rename request for file", c.Fileinfo[0], "in workspace", c.Workspace, "to", c.Fileinfo[1])

	b, err := json.Marshal(c.Workspace)
	utils.CheckError(err)

	ok, err := redisClient.Exists(string(b)).Result()
	utils.CheckError(err)

	if ok == 0 {
		err = getFilesFromDisk(b)
		utils.CheckError(err)
	}

	b1, err := json.Marshal(c.Fileinfo[0])
	utils.CheckError(err)
	b2, err := json.Marshal(c.Fileinfo[1])
	utils.CheckError(err)

	exists, err := redisClient.HExists(string(b), string(b1)).Result()
	utils.CheckError(err)

	if exists {
		data, err := redisClient.HGet(string(b), string(b1)).Result()
		utils.CheckError(err)

		err = redisClient.HDel(string(b), string(b1)).Err()
		utils.CheckError(err)

		err = redisClient.HSet(string(b), string(b2), data).Err()
		utils.CheckError(err)
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.OneFileContainer
	json.Unmarshal(body, &c)

	log.Println("Delete request for file", c.Fileinfo, "in workspace", c.Workspace)

	b1, err := json.Marshal(c.Workspace)
	utils.CheckError(err)

	ok, err := redisClient.Exists(string(b1)).Result()
	utils.CheckError(err)

	if ok == 0 {
		err = getFilesFromDisk(b1)
		utils.CheckError(err)
	}

	b2, err := json.Marshal(c.Fileinfo)
	utils.CheckError(err)

	exists, err := redisClient.HExists(string(b1), string(b2)).Result()
	utils.CheckError(err)

	if exists {
		if !c.Fileinfo.IsDir {
			err = redisClient.HDel(string(b1), string(b2)).Err()
			utils.CheckError(err)
		} else {
			files, err := redisClient.HGetAll(string(b1)).Result()
			utils.CheckError(err)

			for k := range files {
				var fi fileinfo.Fileinfo
				json.Unmarshal([]byte(k), &fi)

				subfile := true

				if len(c.Fileinfo.Path) <= len(fi.Path) {
					for i, v := range c.Fileinfo.Path {
						if v != fi.Path[i] {
							subfile = false
							break
						}
					}
				} else {
					subfile = false
				}

				if subfile {
					b, err := json.Marshal(fi)
					utils.CheckError(err)

					err = redisClient.HDel(string(b1), string(b)).Err()
					utils.CheckError(err)
				}
			}
		}
	}
}

func commitFiles(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Commit request for workspace", string(body))

	bytes, err := getFilesFromCache(body)
	utils.CheckError(err)

	response, err := utils.MakeRequest("http://localhost:10000/api/update", "POST", bytes)
	utils.ForwardResponse(w, response, err)
}

func updateFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.ChangeContainer
	json.Unmarshal(body, &c)

	log.Println("Change request for file", c.Fileinfo, "in workspace", c.Workspace)

	b1, err := json.Marshal(c.Workspace)
	utils.CheckError(err)

	ok, err := redisClient.Exists(string(b1)).Result()
	utils.CheckError(err)

	if ok == 0 {
		err = getFilesFromDisk(b1)
		utils.CheckError(err)
	}

	b2, err := json.Marshal(c.Fileinfo)
	utils.CheckError(err)

	exists, err := redisClient.HExists(string(b1), string(b2)).Result()
	utils.CheckError(err)

	if exists {
		data, err := redisClient.HGet(string(b1), string(b2)).Result()
		utils.CheckError(err)

		data = utils.ApplyChange(data, c.Change)

		err = redisClient.HSet(string(b1), string(b2), data).Err()
		utils.CheckError(err)
	}
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	r := mux.NewRouter()

	r.HandleFunc("/api/workspace", getWorkspace).Methods("POST")
	r.HandleFunc("/api/filetree", getFiletree).Methods("POST")
	r.HandleFunc("/api/clear", clearWorkspace).Methods("POST")

	r.HandleFunc("/api/get", getFile).Methods("POST")
	r.HandleFunc("/api/create", createFile).Methods("POST")
	r.HandleFunc("/api/rename", renameFile).Methods("POST")
	r.HandleFunc("/api/delete", deleteFile).Methods("POST")
	r.HandleFunc("/api/commit", commitFiles).Methods("POST")
	r.HandleFunc("/api/update", updateFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
