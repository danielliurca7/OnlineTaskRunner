package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"

	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"

	//"github.com/rs/xid"
	"../data_structures/containers"
	"../data_structures/workspace"
	"../utils"
)

const tmp = "D:\\Projects\\OnlineTaskRunner\\server-side\\tmp"

// Compile files if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var workspace workspace.Workspace
	json.Unmarshal(body, &workspace)

	dir := utils.GetWorkspaceHash(workspace)
	path := filepath.Join(tmp, dir)

	log.Println("Build request for workspace", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		utils.GetWorkspaceFiles(body, path)
	}

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
	body := utils.GetRequestBody(r)

	var workspace workspace.Workspace
	json.Unmarshal(body, &workspace)

	dir := utils.GetWorkspaceHash(workspace)
	path := filepath.Join(tmp, dir)

	log.Println("Run request for workspace", path)

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
	body := utils.GetRequestBody(r)

	var workspace workspace.Workspace
	json.Unmarshal(body, &workspace)

	dir := utils.GetWorkspaceHash(workspace)
	path := filepath.Join(tmp, dir)

	log.Println("Clean request for workspace", path)

	// Format the make clean command
	cmd := exec.Command("make", "clean")
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = path

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Clean failed with", err)
		utils.WriteResponse(w, 400, e.Bytes())
	} else {
		utils.WriteResponse(w, 200, o.Bytes())
	}
}

// Deletes the workspace folder
func clearRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var workspace workspace.Workspace
	json.Unmarshal(body, &workspace)

	dir := utils.GetWorkspaceHash(workspace)
	path := filepath.Join(tmp, dir)

	err := os.RemoveAll(path)

	if err != nil {
		log.Println("Run failed with", err)
		utils.WriteResponse(w, 400, []byte("Workspace could not be deleted"))
	} else {
		utils.WriteResponse(w, 200, []byte("Workspace deleted"))
	}
}

// Run a specific test
func registerChange(c containers.OneChangeContainer) {
	workspace := utils.GetWorkspace(c.Fileinfo.Workspace)
	path := filepath.Join(tmp, workspace)

	file := filepath.Join(c.Fileinfo.Path...)

	filename := filepath.Join(path, file)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}

	var start, end int64
	var last []byte

	start = c.Change.Position
	end = c.Change.Position + int64(len(c.Change.Previous))

	if int64(len(data)) >= end {
		last = make([]byte, len(data[end:]))
		copy(last, data[end:])
	}

	data = append(append(data[:start], []byte(c.Change.Current)...), last...)

	//log.Println(c.Change, c.Fileinfo)

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		log.Println(err)
	}
}

// Get the file to run from io microservice
func getFile() {

}

func main() {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 8000, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Println(err)
	}

	defer c.Close()

	err = c.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		c.Emit("subscribe", "test.c")
	})
	if err != nil {
		log.Println(err)
	}

	err = c.On("change", func(h *gosocketio.Channel, c containers.OneChangeContainer) {
		registerChange(c)
	})
	if err != nil {
		log.Println(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/build", buildRequest).Methods("PUT")
	r.HandleFunc("/api/run", runRequest).Methods("GET")
	r.HandleFunc("/api/clean", cleanRequest).Methods("DELETE")
	r.HandleFunc("/api/clear", clearRequest).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
