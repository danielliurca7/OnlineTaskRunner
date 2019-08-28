package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"../data_structures/changes"
	"../data_structures/file"
	"../data_structures/fileinfo"
	"../data_structures/tree"
	"../data_structures/workspace"
	"../utils"
	"github.com/gorilla/mux"
)

const fileSystem = "D:\\Projects\\OnlineTaskRunner\\server-side\\file_system"
const sep = string(os.PathSeparator)

// Creates the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var fi fileinfo.Fileinfo
	json.Unmarshal(body, &fi)

	path := utils.GetPath(fi)

	log.Println("Create file request for path " + path)

	dir, file := filepath.Split(path)

	path = filepath.Join(fileSystem, dir)

	if !utils.FileExists(path) {
		err := os.MkdirAll(path, os.ModeDir)

		if err != nil {
			log.Println(err)
			utils.WriteResponse(w, 500, []byte("Folder could not be created"))
			return
		}
	}

	path = filepath.Join(path, file)

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

	var fi []fileinfo.Fileinfo
	json.Unmarshal(body, &fi)

	oldPath := utils.GetPath(fi[0])
	newPath := utils.GetPath(fi[1])

	log.Println("Rename file request for path " + oldPath)

	if !utils.FileExists(oldPath) {
		utils.WriteResponse(w, 404, []byte("File to rename not found"))
	} else {
		err := os.Rename(oldPath, newPath)

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

	var fi fileinfo.Fileinfo
	json.Unmarshal(body, &fi)

	path := utils.GetPath(fi)

	log.Println("Delete file request for path", path)

	path = filepath.Join(fileSystem, path)

	log.Println(fi)

	if !utils.FileExists(path) {
		utils.WriteResponse(w, 400, []byte("No file at the specified path"))
	} else {
		err := os.Remove(path)

		if err != nil {
			log.Println(err)
			utils.WriteResponse(w, 500, []byte("File could not be deleted"))
			return
		}

		utils.WriteResponse(w, 200, []byte("File deleted"))
	}
}

// Establishes a tcp connection and sends a file
func getFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var fi fileinfo.Fileinfo
	json.Unmarshal(body, &fi)

	path := utils.GetPath(fi)

	log.Println("Get file request for path", path)

	path = filepath.Join(fileSystem, path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		utils.WriteResponse(w, 500, []byte("File could not be read"))
		log.Println(err)
	}

	utils.WriteResponse(w, 200, data)
}

func getFileTree(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var ws workspace.Workspace
	json.Unmarshal(body, &ws)

	path := utils.GetWorkspace(ws)

	log.Println("Get file tree request for path", path)

	path = filepath.Join(fileSystem, path)
	pathLen := len(strings.Split(path, sep))

	root := &tree.Node{}

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			pathParts := strings.Split(path, sep)[pathLen:]

			presentNode := root

			for index, fileName := range pathParts {
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
						IsDir:    info.IsDir(),
						Path:     pathParts[:index+1],
					}
					presentNode.Children = append(presentNode.Children, node)
					presentNode = node
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
		utils.WriteResponse(w, 500, []byte("Could not walk through file tree"))
	}

	b, err := json.Marshal(root.Children)
	utils.CheckError(err)

	utils.WriteResponse(w, 200, b)
}

func getWorkSpace(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var ws workspace.Workspace
	json.Unmarshal(body, &ws)

	path := utils.GetWorkspace(ws)

	log.Println("Get file tree request for path", path)

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

			path = strings.Join(
				strings.Split(
					path, sep,
				)[pathLen:],
				sep,
			)

			f := file.File{
				Path:  path,
				Data:  string(data),
				IsDir: info.IsDir(),
			}

			if path != "" {
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

	var changesList changes.Changes
	json.Unmarshal(body, &changesList)

	log.Println("Update file request for", utils.GetPath(changesList.Fileinfo))

	path := filepath.Join(fileSystem, utils.GetPath(changesList.Fileinfo))

	data, err := ioutil.ReadFile(path)
	utils.CheckError(err)

	for _, change := range changesList.Changes {
		data = utils.ApplyChange(data, change)
	}

	err = ioutil.WriteFile(path, data, 0666)
	utils.CheckError(err)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/file", createFile).Methods("POST")
	r.HandleFunc("/api/file", renameFile).Methods("PUT")
	r.HandleFunc("/api/file", deleteFile).Methods("DELETE")
	r.HandleFunc("/api/get", getFile).Methods("POST")
	r.HandleFunc("/api/file", updateFiles).Methods("PATCH")

	r.HandleFunc("/api/filetree", getFileTree).Methods("GET")
	r.HandleFunc("/api/workspace", getWorkSpace).Methods("GET")

	log.Fatal(http.ListenAndServe(":8002", r))
}
