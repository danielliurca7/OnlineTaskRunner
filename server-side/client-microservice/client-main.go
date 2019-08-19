package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"../data_structures/change"
	"../data_structures/changes"
	"../data_structures/containers"
	"../utils"
	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

const bufSize = 5

// Changes is a
var Changes []changes.Changes

// Sends request to create the file with the specified name
func createFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Create file request")

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:8002/api/file", "POST", body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		utils.WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

// Sends request to rename the file with the specified name
func renameFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Rename file request")

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:8002/api/file", "PUT", body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utils.WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

// Sends request to delete the file with the specified name
func deleteFile(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Delete file request")

	// Make a request to the io microservice
	response, err := utils.MakeRequest("http://localhost:8002/api/file", "DELETE", body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utils.WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

// Register if the client wrote some code
// Update the buffer and notify observers
func registerChange(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	var c containers.OneChangeContainer
	json.Unmarshal(body, &c)

	ch.Emit("change", c)

	i, changeList := utils.GetFileChanges(Changes, c.Fileinfo)

	if i == -1 {
		Changes = append(Changes, changes.Changes{
			Fileinfo: c.Fileinfo,
			Changes:  []change.Change{c.Change},
		})
	} else {
		changeList = append(changeList, c.Change)

		if len(changeList) >= bufSize {
			//commit changes to io microservice
			changeList = nil
		}

		Changes = append(
			append(Changes[:i], changes.Changes{
				Fileinfo: c.Fileinfo,
				Changes:  changeList,
			},
			),
			Changes[i+1:]...,
		)
	}
}

// Send changes to obervers
func commitChanges() {

}

// If the client wants to download the file
// The client microservice verify that he has permissions for the file
// And gives the client a port to listen to for the io microservice
func verifyConnection(w http.ResponseWriter, r *http.Request) {

}

// Make a request to the compute microservice
// To compile a workspace if necessary
func buildRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Build request for workspace " + string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://localhost:8001/api/build", "PUT", body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utils.WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

// Make a request to the compute microservice
// To run an executable/interpretable code
func runRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Run request for workspace " + string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://localhost:8001/api/run", "GET", body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utils.WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

// Make a request to the compute microservice
// To clean the workspace
func cleanRequest(w http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody(r)

	log.Println("Clean request for workspace " + string(body))

	// Make a request to the compute microservice
	response, err := utils.MakeRequest("http://localhost:8001/api/clean", "DELETE", body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utils.WriteResponse(w, response.StatusCode, data)
	}

	defer response.Body.Close()
}

// Creates a custom test for a certain workspace
func createCustomTest(w http.ResponseWriter, r *http.Request) {

}

// Edits the build file
func editBuild(w http.ResponseWriter, r *http.Request) {

}

// Edits a task file
func editTasks(w http.ResponseWriter, r *http.Request) {

}

// Edits the docker image
func editDockerImage(w http.ResponseWriter, r *http.Request) {

}

var ch *gosocketio.Channel

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("Connected")
		ch = c
	})

	server.On("subscribe", func(c *gosocketio.Channel, s string) {
		c.Join(s)
	})

	server.On("unsubscribe", func(c *gosocketio.Channel, s string) {
		c.Leave(s)
	})

	server.On("change", func(h *gosocketio.Channel, c containers.OneChangeContainer) {
		h.BroadcastTo(filepath.Join(c.Fileinfo.Path...), "change", c)
	})

	r := mux.NewRouter()

	r.Handle("/socket.io/", server)

	r.HandleFunc("/api/file", createFile).Methods("POST")
	r.HandleFunc("/api/file", renameFile).Methods("PUT")
	r.HandleFunc("/api/file", deleteFile).Methods("DELETE")
	r.HandleFunc("/api/file", verifyConnection).Methods("GET")

	r.HandleFunc("/api/change", registerChange).Methods("PUT")

	r.HandleFunc("/api/request", buildRequest).Methods("PUT")
	r.HandleFunc("/api/request", runRequest).Methods("GET")
	r.HandleFunc("/api/request", cleanRequest).Methods("DELETE")

	r.HandleFunc("/api/test", createCustomTest).Methods("POST")
	r.HandleFunc("/api/build", editBuild).Methods("PUT")
	r.HandleFunc("/api/tasks", editTasks).Methods("PUT")
	r.HandleFunc("/api/image", editDockerImage).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
