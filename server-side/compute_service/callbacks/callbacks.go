package callbacks

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

	_, cmdErr, err := utils.RunCommand(
		image.Workspace.ToString(),
		"docker", "build", ".", "-t", imageName, "--no-cache",
	)

	if err != nil {
		log.Println("Build failed with", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Build failed"))
	} else if len(cmdErr) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(append([]byte("Build failed"), cmdErr...))
	} else {
		w.Write([]byte("Build successful"))
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
	h = sha1.New()
	h.Write([]byte(ws.ToString()))
	instanceName := string(hex.EncodeToString(h.Sum(nil)))

	files = append(files, datastructures.GetDockerFile(imageName))

	if err := utils.CopyFilesToMemory(ws, files); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get copy files into memory"))
		return
	}

	if _, _, err = utils.RunCommand(
		ws.ToString(),
		"docker", "run", "--rm", "-t", "-d", "--name", instanceName, imageName,
	); err != nil {
		log.Println(err)
	}

	if _, _, err := utils.RunCommand(
		ws.ToString(),
		"docker", "cp", ".", fmt.Sprintf("%s:/app", instanceName),
	); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not copy file"))
		return
	}

	cmdOut, cmdErr, err := utils.RunCommand(
		ws.ToString(),
		"docker", "exec", instanceName, "./run.sh",
	)

	// Run the command
	if err != nil {
		log.Println("Run failed with", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Run failed\n"))
	} else if len(cmdErr) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(append([]byte("Run failed\n"), cmdErr...))
	} else {
		w.Write(cmdOut)
	}
}

// StopContainer is the callback for a stop request
// It executes a docker stop command
func StopContainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not read request body"))
		return
	}

	var ws datastructures.Workspace

	if err := json.Unmarshal(body, &ws); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not read request body"))
		return
	}

	h := sha1.New()
	h.Write([]byte(ws.ToString()))
	instanceName := string(hex.EncodeToString(h.Sum(nil)))

	_, cmdErr, err := utils.RunCommand(
		ws.ToString(),
		"docker", "stop", instanceName,
	)

	if err != nil {
		log.Println("Stop failed with", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Stop failed"))
	} else if len(cmdErr) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(append([]byte("Stop failed\n"), cmdErr...))
	}
}
