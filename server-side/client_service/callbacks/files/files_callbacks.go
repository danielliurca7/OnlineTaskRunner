package files

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/streadway/amqp"

	datastructures "../../data_structures"
	"../../utils"
)

// GetFiles verifies the the validity of the token and the request body
// If token is valid, redirects the request to files service
func GetFiles(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var workspace datastructures.Workspace
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &workspace); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("GET", "http://files:4000/api/files", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)
}

// CommitFiles verifies the the validity of the token and the request body
// If token is valid, redirects the request to files service
func CommitFiles(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var workspace datastructures.Workspace
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &workspace); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("POST", "http://files:4000/api/commit", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)
}

// CreateFile verifies the the validity of the token and the request body
// If token is valid, redirects the request to files service
func CreateFile(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var createBody datastructures.CreateBody
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &createBody); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(createBody.Workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("POST", "http://files:4000/api/create", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)
}

// DeleteFile verifies the the validity of the token and the request body
// If token is valid, redirects the request to files service
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var deleteBody datastructures.DeleteBody
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &deleteBody); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(deleteBody.Workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("DELETE", "http://files:4000/api/delete", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)
}

// RenameFile verifies the the validity of the token and the request body
// If token is valid, redirects the request to files service
func RenameFile(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	token := r.Header.Get("Authorization")
	var renameBody datastructures.RenameBody
	var userdata datastructures.UserData

	if err := json.Unmarshal(body, &renameBody); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userdata, err = utils.VerifyToken(token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.VerifyAuthorization(renameBody.Workspace, userdata) {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, status, err := utils.MakeRequest("PUT", "http://files:4000/api/rename", body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(data)
}

// ListenForUpdates listens to a RabbitMQ broker for updates for files
// When an update occurs send it to the client
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
			"Update request for",
			"workspace", body.Workspace.ToString(),
			"path", filepath.Join(body.Path...),
			"change", body.Change,
		)
	}
}
