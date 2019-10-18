package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../data_structures/containers"
	"../utils"
)

func get(bytes []byte) containers.WorkspaceContainer {
	response, err := utils.MakeRequest("http://localhost:7000/api/workspace", "POST", bytes)
	utils.CheckError(err)

	bytes = utils.GetResponseBody(response)

	var c containers.WorkspaceContainer
	json.Unmarshal(bytes, &c)

	return c
}

func build(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	files := get(body)

	imageName := utils.GetPathFromBody(body)
	p := filepath.Join(utils.Tmp, imageName)

	err := os.MkdirAll(p, 0666)
	utils.CheckError(err)

	for _, v := range files.Files {
		path := filepath.Join(p, filepath.Join(v.Path...))

		if !v.IsDir {
			err := ioutil.WriteFile(path, []byte(v.Data), 0666)
			utils.CheckError(err)
		} else {
			err := os.MkdirAll(path, 0666)
			utils.CheckError(err)
		}
	}

	log.Println("Build request for workspace", imageName)

	// Format the make build command
	//cmd := exec.Command("/bin/sh", "-c", "docker build", ".", "-t", imageName)
	cmd := exec.Command("cmd", "/C", "docker build", ".", "-t", imageName)
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = p

	// Run the command
	err = cmd.Run()
	if err != nil {
		log.Println("Build failed with", err)
		utils.WriteResponse(w, 400, append([]byte("Build\n"), e.Bytes()...))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			utils.WriteResponse(w, 200, []byte("Build successful"))
		} else {
			utils.WriteResponse(w, 200, append([]byte("Build\n"), e.Bytes()...))
		}
	}
}

func run(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	imageName := utils.GetPathFromBody(body)

	log.Println("Run request for workspace", imageName)

	// Format the make build command
	//cmd := exec.Command("/bin/sh", "-c", "docker run -d", "--name", imageName, imageName)
	cmd := exec.Command("cmd", "/C", "docker run", "--name", imageName, imageName)
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Build failed with", err)
		utils.WriteResponse(w, 400, append([]byte("Run\n"), e.Bytes()...))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			utils.WriteResponse(w, 200, append([]byte("Run\n"), o.Bytes()...))
		} else {
			utils.WriteResponse(w, 200, append([]byte("Run\n"), e.Bytes()...))
		}
	}

	clear(imageName)
}

func stop(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	imageName := utils.GetPathFromBody(body)
	path := utils.Tmp

	log.Println("Stop request for workspace", imageName)

	// Format the make build command
	cmd := exec.Command("/bin/sh", "-c",
		"docker stop", imageName, ";",
		"docker rm", imageName, ";",
	)
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = path

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Stop failed with", err)
		utils.WriteResponse(w, 400, append([]byte("Stop\n"), e.Bytes()...))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			utils.WriteResponse(w, 200, []byte("Stop successful"))
		} else {
			utils.WriteResponse(w, 200, append([]byte("Stop\n"), e.Bytes()...))
		}
	}
}

func clear(imageName string) {
	path := utils.Tmp

	log.Println("Clear request for workspace", imageName)

	// Format the make build command
	//cmd := exec.Command("/bin/sh", "-c", "rm -fr", imageName)
	/*cmd := exec.Command(
		"/bin/sh", "-c", "docker rm", imageName, ";",
		"/bin/sh", "-c", "del /fr", imageName,
	)*/
	cmd := exec.Command(
		"cmd", "/C", "docker rm", imageName,
	)
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	cmd.Dir = path

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Clear failed with", err)
		log.Println(string(append([]byte("Clear"), e.Bytes()...)))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			log.Println("Clear successful")
		} else {
			log.Println(string(append([]byte("Clear\n"), e.Bytes()...)))
		}
	}
}

func main() {
	log.Println("Computer microservice is running")

	r := mux.NewRouter()

	r.HandleFunc("/api/build", build).Methods("POST")
	r.HandleFunc("/api/run", run).Methods("POST")
	r.HandleFunc("/api/stop", stop).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
