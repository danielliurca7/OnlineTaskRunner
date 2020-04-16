package callbacks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	datastructures "../data_structures"
	"../utils"
	"github.com/streadway/amqp"
)

// GetFiles is the callback that returns files in the workspace specified in the request
// Files are return from cache, if they are there, or from the disk
// If files are not in cache we will copy them there
func GetFiles(w http.ResponseWriter, r *http.Request) {
	var ws datastructures.Workspace

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not read request body"))
		return
	} else if err = json.Unmarshal(body, &ws); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	log.Println("Get files request for workspace", ws.ToString())

	var files []datastructures.File

	filesInCache, err := utils.IsWorkspaceInCache(ws)

	if filesInCache {
		files, err = utils.GetFilesFromCache(ws)
	} else {
		files, err = utils.GetFilesFromDisk(ws)

		copy, ok := r.URL.Query()["copy"]

		if ok && copy[0] == "YES" {
			go func() {
				if err = utils.GetFilesToCache(ws, files); err != nil {
					log.Println(err)
				}
			}()
		}
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Files could not be read"))
		return
	}

	if b, err := json.Marshal(files); err != nil {
		log.Println(err)
	} else {
		w.Write(b)
	}
}

// CommitFiles is the callback that commits files in the workspace specified in the request to disk
func CommitFiles(w http.ResponseWriter, r *http.Request) {
	var ws datastructures.Workspace

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not read request body"))
		return
	} else if err = json.Unmarshal(body, &ws); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	log.Println("Commit files request for workspace", ws.ToString())

	if err := utils.CommitFilesToDisk(ws); err != nil {
		log.Println(err)
	}

	if err := utils.ClearFilesFromCache(ws); err != nil {
		log.Println(err)
	}
}

// CreateFile creates a new file in a workspace
func CreateFile(w http.ResponseWriter, r *http.Request) {
	var body datastructures.CreateBody

	if reqBody, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not read request body"))
		return
	} else if err = json.Unmarshal(reqBody, &body); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	log.Println(
		"Create file request for",
		"workspace", body.Workspace.ToString(),
		"path", filepath.Join(body.File.Path...),
	)

	if err := utils.CreateFile(body.Workspace, body.File); err != nil {
		log.Println(err)
	}
}

// DeleteFile deletes a file from a workspace
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	var body datastructures.DeleteBody

	if reqBody, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not read request body"))
		return
	} else if err = json.Unmarshal(reqBody, &body); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	log.Println(
		"Delete file request for",
		"workspace", body.Workspace.ToString(),
		"path", filepath.Join(body.Path...),
	)

	if err := utils.DeleteFile(body.Workspace, body.Path); err != nil {
		log.Println(err)
	}
}

// RenameFile renames a file in a workspace to a newName
func RenameFile(w http.ResponseWriter, r *http.Request) {
	var body datastructures.RenameBody

	if reqBody, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Could not read request body"))
		return
	} else if err = json.Unmarshal(reqBody, &body); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	log.Println(
		"Rename file request for",
		"workspace", body.Workspace.ToString(),
		"path", filepath.Join(body.Path...),
	)

	if err := utils.RenameFile(body.Workspace, body.Path, body.NewName); err != nil {
		log.Println(err)
	}
}

// ListenForUpdates listens to a RabbitMQ broker for updates for files
// Updates can occur only if files are in cache(i.e. someone requested them first)
func ListenForUpdates() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer ch.Close()

	if err = ch.ExchangeDeclare(
		"changes", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	); err != nil {
		log.Fatalln(err)
		return
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if err = ch.QueueBind(
		q.Name,    // queue name
		"",        // routing key
		"changes", // exchange
		false,
		nil,
	); err != nil {
		log.Fatalln(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalln(err)
		return
	}

	for msg := range msgs {
		var body datastructures.UpdateBody

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			log.Println(err)
		}

		log.Println(
			"Update file request for",
			"workspace", body.Workspace.ToString(),
			"path", filepath.Join(body.Path...),
			"change", body.Change,
		)

		if err := utils.UpdateFile(body.Workspace, body.Path, body.Change); err != nil {
			log.Println(err)
		}
	}
}
