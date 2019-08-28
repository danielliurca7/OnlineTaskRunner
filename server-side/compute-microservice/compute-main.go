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

	"../data_structures/containers"
	"../utils"
)

func get(w http.ResponseWriter, body []byte) {
	info := utils.GetFileinfoFromBody(body)

	b, err := json.Marshal(info.Workspace)
	utils.CheckError(err)

	dir := utils.GetWorkspaceHash(info.Workspace)
	path := filepath.Join(utils.Tmp, dir)

	log.Println("Get request for workspace", path)

	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = utils.GetWorkspaceFiles(b, path)
	}

	if err != nil {
		log.Println("Get failed with", err)
		utils.WriteResponse(w, 400, []byte("Workspace could not be copied"))
	} else {
		utils.WriteResponse(w, 200, []byte("Workspace copied"))
	}

	ch.Emit("subscribe", info)
}

func build(w http.ResponseWriter, body []byte) {
	path := utils.GetPathFromBody(body)

	log.Println("Build request for workspace", path)

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
		utils.WriteResponse(w, 400, append([]byte("Build\n"), e.Bytes()...))
	} else {
		utils.WriteResponse(w, 200, append([]byte("Build\n"), e.Bytes()...))
	}
}

func run(w http.ResponseWriter, body []byte) {
	path := utils.GetPathFromBody(body)

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
		utils.WriteResponse(w, 400, append([]byte("Run\n"), e.Bytes()...))
	} else {
		utils.WriteResponse(w, 200, append([]byte("Run\n"), o.Bytes()...))
	}
}

func clean(w http.ResponseWriter, body []byte) {
	path := utils.GetPathFromBody(body)

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
		utils.WriteResponse(w, 400, append([]byte("Clean\n"), e.Bytes()...))
	} else {
		utils.WriteResponse(w, 200, append([]byte("Clean\n"), o.Bytes()...))
	}
}

func clear(w http.ResponseWriter, body []byte) {
	path := utils.GetPathFromBody(body)

	err := os.RemoveAll(path)

	log.Println("Clear request for workspace", path)

	if err != nil {
		log.Println("Run failed with", err)
		utils.WriteResponse(w, 400, []byte("Workspace could not be deleted"))
	} else {
		utils.WriteResponse(w, 200, []byte("Workspace deleted"))
	}
}

// Get the workspace from io microservice
func getRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	get(w, body)
}

// Compile files if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	build(w, body)
}

// Runs the project in a workspace
func runRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	run(w, body)
}

// Cleans the project in a workspace
func cleanRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	clean(w, body)
}

// Deletes the workspace folder
func clearRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	clear(w, body)
}

// Run a specific test
func registerChange(c containers.OneChangeContainer) {
	workspace := utils.GetWorkspaceHash(c.Fileinfo.Workspace)
	path := filepath.Join(utils.Tmp, workspace)

	file := filepath.Join(c.Fileinfo.Path...)

	filename := filepath.Join(path, file)

	log.Println("Change on file", filename)

	data, err := ioutil.ReadFile(filename)
	utils.CheckError(err)

	data = utils.ApplyChange(data, c.Change)

	err = ioutil.WriteFile(filename, data, 0666)
	utils.CheckError(err)
}

var ch *gosocketio.Channel

func main() {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 8000, false),
		transport.GetDefaultWebsocketTransport())
	utils.CheckError(err)

	defer c.Close()

	err = c.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		ch = c
	})
	utils.CheckError(err)

	err = c.On("change", func(h *gosocketio.Channel, c containers.OneChangeContainer) {
		registerChange(c)
	})
	utils.CheckError(err)

	r := mux.NewRouter()

	r.HandleFunc("/api/get", getRequest).Methods("GET")
	r.HandleFunc("/api/build", buildRequest).Methods("PUT")
	r.HandleFunc("/api/run", runRequest).Methods("GET")
	r.HandleFunc("/api/clean", cleanRequest).Methods("DELETE")
	r.HandleFunc("/api/clear", clearRequest).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
