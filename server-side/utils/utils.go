package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"../data_structures/change"
	"../data_structures/changes"
	"../data_structures/file"
	"../data_structures/fileinfo"
	"../data_structures/workspace"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func GetRequestBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Couldn't read request body")
	}

	return body
}

func GetResponseBody(r *http.Response) []byte {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Couldn't read response body")
	}

	return body
}

func MakeRequest(host string, method string, value []byte) (*http.Response, error) {
	request, _ := http.NewRequest(method, host, bytes.NewBuffer(value))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(request)
}

func WriteResponse(w http.ResponseWriter, statusCode int, data []byte) {
	w.WriteHeader(statusCode)
	w.Write(data)
}

func GetFileChanges(changes []changes.Changes, fileinfo fileinfo.Fileinfo) (int, []change.Change) {
	for i, c := range changes {
		if c.Fileinfo.Equals(&fileinfo) {
			return i, c.Changes
		}
	}

	return -1, nil
}

func GetWorkspace(w workspace.Workspace) string {
	path := filepath.Join(w.Subject, strconv.Itoa(w.Year), w.AssignmentName, w.Owner)

	return path
}

func GetWorkspaceHash(w workspace.Workspace) string {
	path := filepath.Join(w.Subject, strconv.Itoa(w.Year), w.AssignmentName, w.Owner)
	h := sha1.New()
	h.Write([]byte(path))
	bytes := hex.EncodeToString(h.Sum(nil))

	return string(bytes)
}

func GetPath(fi fileinfo.Fileinfo) string {
	extPath := filepath.Join(fi.Subject, strconv.Itoa(fi.Year), fi.AssignmentName, fi.Owner)
	intPath := filepath.Join(fi.Path...)

	return filepath.Join(extPath, intPath)
}

func GetWorkspaceFiles(workspace []byte, path string) {
	response, err := MakeRequest("http://localhost:8002/api/workspace", "GET", workspace)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return
	}

	body := GetResponseBody(response)

	var files []file.File

	json.Unmarshal(body, &files)

	err = os.Mkdir(path, 0666)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		filepath := filepath.Join(path, file.Path)

		if file.IsDir {
			err = os.Mkdir(filepath, 0666)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = ioutil.WriteFile(filepath, []byte(file.Data), 0666)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
