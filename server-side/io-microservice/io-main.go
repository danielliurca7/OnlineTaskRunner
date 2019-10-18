package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"../data_structures/containers"
	"../data_structures/file"
	"../data_structures/workspace"
	"../utils"
	"github.com/gorilla/mux"
)

const fileSystem = "D:\\Projects\\OnlineTaskRunner\\server-side\\file_system"
const sep = string(os.PathSeparator)

func getWorkSpace(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var ws workspace.Workspace
	json.Unmarshal(body, &ws)

	path := utils.GetWorkspace(ws)

	log.Println("Get workspace request for path", path)

	path = filepath.Join(fileSystem, path)
	pathLen := len(strings.Split(path, sep))

	var files []file.File

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			var data []byte

			if !info.IsDir() {
				data, err = ioutil.ReadFile(path)
				if err != nil {
					utils.WriteResponse(w, 500, []byte("A File could not be read"))
					log.Println(err)
				}
			} else {
				data = nil
			}

			pathList := strings.Split(path, sep)[pathLen:]

			f := file.File{
				Path:  pathList,
				Data:  string(data),
				IsDir: info.IsDir(),
			}

			if len(pathList) > 0 {
				files = append(files, f)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
		utils.WriteResponse(w, 500, []byte("Could not walk through file tree"))
	}

	b, err := json.Marshal(files)
	utils.CheckError(err)

	utils.WriteResponse(w, 200, b)
}

// Modifies and stores the file according to the request
func updateFiles(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.WorkspaceContainer
	json.Unmarshal(body, &c)

	workspace := filepath.Join(fileSystem, utils.GetWorkspace(c.Workspace))

	log.Println("Update workspace request for", workspace)

	dir, err := ioutil.ReadDir(workspace)
	utils.CheckError(err)

	for _, d := range dir {
		os.RemoveAll(path.Join([]string{workspace, d.Name()}...))
	}

	for _, v := range c.Files {
		path := filepath.Join(workspace, filepath.Join(v.Path...))

		if !v.IsDir {
			err := ioutil.WriteFile(path, []byte(v.Data), 0666)
			utils.CheckError(err)
		} else {
			err := os.MkdirAll(path, 0666)
			utils.CheckError(err)
		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/update", updateFiles).Methods("POST")
	r.HandleFunc("/api/workspace", getWorkSpace).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", r))
}
