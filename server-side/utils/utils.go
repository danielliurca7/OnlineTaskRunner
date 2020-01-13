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
	"strings"

	"../data_structures/change"
	"../data_structures/changes"
	"../data_structures/file"
	"../data_structures/fileinfo"
	"../data_structures/workspace"
)

//const Tmp = "/tmp"
const Tmp = "D:\\Projects\\OnlineTaskRunner\\server-side\\tmp"
const sep = string(os.PathSeparator)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

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

func ForwardResponse(w http.ResponseWriter, response *http.Response, err error) {
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

func WriteResponse(w http.ResponseWriter, statusCode int, data []byte) {
	w.WriteHeader(statusCode)
	w.Write(data)
}

func GetFileChanges(c []changes.Changes, fileinfo fileinfo.Fileinfo) (int, changes.Changes) {
	for i, change := range c {
		if change.Fileinfo.Equals(&fileinfo) {
			return i, change
		}
	}

	return -1, changes.Changes{}
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

func GetPathFromBody(body []byte) string {
	var workspace workspace.Workspace
	json.Unmarshal(body, &workspace)

	dir := GetWorkspaceHash(workspace)

	return dir
}

func GetFileinfoFromBody(body []byte) fileinfo.Fileinfo {
	var fileinfo fileinfo.Fileinfo
	json.Unmarshal(body, &fileinfo)

	return fileinfo
}

func GetWorkspaceFiles(workspace []byte, path string) error {
	response, err := MakeRequest("http://localhost:10000/api/workspace", "POST", workspace)

	if err != nil {
		return err
	}

	body := GetResponseBody(response)

	var files []file.File

	json.Unmarshal(body, &files)

	err = os.Mkdir(path, 0666)
	if err != nil {
		return err
	}

	for _, file := range files {
		filepath := filepath.Join(path, filepath.Join(file.Path...))

		if file.IsDir {
			err = os.Mkdir(filepath, 0666)
			if err != nil {
				return err
			}
		} else {
			err = ioutil.WriteFile(filepath, []byte(file.Data), 0666)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ApplyChange(content string, c change.Change) string {
	data := []byte(content)

	var start, end int64
	var last []byte

	start = c.Position
	end = c.Position + int64(len(c.Previous))

	if int64(len(data)) >= end {
		last = make([]byte, len(data[end:]))
		copy(last, data[end:])
	}

	return string(append(append(data[:start], []byte(c.Current)...), last...))
}

func ReadFiles(path string, files []file.File) ([]file.File, error) {
	pathLen := len(strings.Split(path, sep))

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			var data []byte

			if !info.IsDir() {
				data, err = ioutil.ReadFile(path)
				if err != nil {
					return err
				}
			} else {
				data = nil
			}

			pathList := strings.Split(path, sep)[pathLen-1:]

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

	return files, err
}
