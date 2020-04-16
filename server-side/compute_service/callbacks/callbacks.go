package callbacks

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	datastructures "../data_structures"
	"../utils"
)

// BuildImage is the callback for a build request
// It executes a docker build command
func BuildImage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	var image datastructures.BuildBody

	if err := json.Unmarshal(body, &image); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	var files []datastructures.File

	if b, err := json.Marshal(image.Workspace); err != nil {
		log.Println(err)
	} else if files, err = utils.GetConfigFiles(b); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not get necessary files"))
		return
	}

	files = append(files, datastructures.GetDockerFile(image.Image))

	if err := utils.CopyFilesToMemory(image.Workspace, files); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not get copy files into memory"))
		return
	}

	h := sha1.New()
	h.Write([]byte(image.Workspace.Name()))
	imageName := string(hex.EncodeToString(h.Sum(nil)))

	cmd := exec.Command("docker", "build", ".", "-t", imageName, "--no-cache")
	var e bytes.Buffer
	cmd.Stderr = &e
	cmd.Dir = image.Workspace.ToString()

	// Run the command
	err = cmd.Run()
	if err != nil {
		log.Println("Build failed with", err)
		w.WriteHeader(400)
		w.Write(append([]byte("Build failed\n"), e.Bytes()...))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			w.Write(append([]byte("Build successful\n"), e.Bytes()...))
		} else {
			w.WriteHeader(400)
			w.Write(append([]byte("Build failed\n"), e.Bytes()...))
		}
	}

}

// RunContainer is the callback for a run request
// It executes a docker run command
func RunContainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	var ws datastructures.Workspace
	var files []datastructures.File

	if files, err = utils.GetConfigFiles(body); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not get necessary files"))
		return
	} else if err = json.Unmarshal(body, &ws); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	h := sha1.New()
	h.Write([]byte(ws.Name()))
	imageName := string(hex.EncodeToString(h.Sum(nil)))
	h.Write([]byte(ws.ToString()))
	instanceName := string(hex.EncodeToString(h.Sum(nil)))

	files = append(files, datastructures.GetDockerFile(imageName))

	if err := utils.CopyFilesToMemory(ws, files); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not get copy files into memory"))
		return
	}

	cmd := exec.Command("docker", "build", ".", "-t", instanceName, "--no-cache")
	var o, e bytes.Buffer
	cmd.Stderr = &e
	cmd.Dir = ws.ToString()

	if err := cmd.Run(); err != nil {
		log.Println("Build failed with", err)
		w.WriteHeader(400)
		w.Write(append([]byte("Run failed\n"), e.Bytes()...))
	}

	cmd = exec.Command("docker", "run", "--rm", instanceName)
	cmd.Stdout = &o
	cmd.Stderr = &e

	// Run the command
	if err = cmd.Run(); err != nil {
		log.Println("Run failed with", err)
		w.WriteHeader(400)
		w.Write(append([]byte("Run failed\n"), e.Bytes()...))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			w.Write(o.Bytes())
		} else {
			w.WriteHeader(400)
			w.Write(append([]byte("Run failed\n"), e.Bytes()...))
		}
	}
}

// StopContainer is the callback for a stop request
// It executes a docker stop command
func StopContainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	var ws datastructures.Workspace

	if err := json.Unmarshal(body, &ws); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Could not read request body"))
		return
	}

	h := sha1.New()
	h.Write([]byte(ws.ToString()))
	instanceName := string(hex.EncodeToString(h.Sum(nil)))

	// Format the make build command
	cmd := exec.Command("docker", "stop", instanceName)
	var e bytes.Buffer
	cmd.Stderr = &e

	// Run the command

	if err := cmd.Run(); err != nil {
		log.Println("Stop failed with", err)
		w.WriteHeader(400)
		w.Write(append([]byte("Stop failed\n"), e.Bytes()...))
	} else {
		if len([]byte(e.Bytes())) == 0 {
			w.Write([]byte("Stop succesful\n"))
		} else {
			w.WriteHeader(400)
			w.Write(append([]byte("Stop failed\n"), e.Bytes()...))
		}
	}
}
