package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

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

	log.Println("Get workspace request for path", filepath.Join(ws.Subject, strconv.Itoa(ws.Year), ws.AssignmentName, ws.Owner))

	var files []file.File

	path := filepath.Join(fileSystem, ws.Subject, strconv.Itoa(ws.Year), ws.AssignmentName, ws.Owner)

	files, err := utils.ReadFiles(path, files)
	if err != nil {
		log.Println(err)
		utils.WriteResponse(w, 500, []byte("Could not walk through file tree"))
	}

	path = filepath.Join(fileSystem, ws.Subject, strconv.Itoa(ws.Year), ws.AssignmentName, "tests")

	files, err = utils.ReadFiles(path, files)
	if err != nil {
		log.Println(err)
		utils.WriteResponse(w, 500, []byte("Could not walk through file tree"))
	}

	path = filepath.Join(fileSystem, ws.Subject, strconv.Itoa(ws.Year), ws.AssignmentName, "libraries")

	files, err = utils.ReadFiles(path, files)
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

	log.Fatal(http.ListenAndServe(":11000", r))
}
