package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// MakeRequest sends a http request to specified url and body
// It return the response data, status and eventual error
func MakeRequest(method string, url string, body []byte) ([]byte, int, error) {
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return data, response.StatusCode, nil
}
