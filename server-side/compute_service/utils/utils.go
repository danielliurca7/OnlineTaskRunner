package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	datastructures "../data_structures"
)

// CopyFilesToMemory copies the files given as paramater to tmp, in a folder dependent on the workspace
func CopyFilesToMemory(ws datastructures.Workspace, files []datastructures.File) error {
	if err := os.MkdirAll(ws.ToString(), 0666); err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(ws.ToString(), filepath.Join(file.Path...))

		if !file.IsDir {
			if err := ioutil.WriteFile(path, []byte(file.Data), 0666); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(path, 0666); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetConfigFiles call the files service and returns the result
func GetConfigFiles(body []byte) ([]datastructures.File, error) {
	request, _ := http.NewRequest("GET", "http://files:4000/api/files", bytes.NewBuffer(body))
	client := &http.Client{}

	if response, err := client.Do(request); err != nil {
		return nil, err
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else {
		var files []datastructures.File

		if err := json.Unmarshal(body, &files); err != nil {
			return nil, err
		}

		return files, nil
	}
}
