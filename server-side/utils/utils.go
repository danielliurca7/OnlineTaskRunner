package utils

import (
	"os"
	"bytes"
	"net/http"
	"io/ioutil"
	"log"
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
