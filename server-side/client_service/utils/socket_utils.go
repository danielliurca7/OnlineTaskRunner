package utils

import (
	"encoding/json"
	"log"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/streadway/amqp"

	datastructures "../data_structures"
)

// NewChangeServer creates a socket server for async file change transactions with the client
func NewChangeServer() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On("subscribe", func(c *gosocketio.Channel, body datastructures.SubscribeBody) {
		var userdata datastructures.UserData
		var err error

		if userdata, err = VerifyToken(body.Token); err != nil {
			log.Println(err)
			c.Close()
			return
		}

		if !VerifyAuthorization(body.Workspace, userdata) {
			log.Println("Unauthorized")
			c.Close()
			return
		}
		c.Join(body.Workspace.ToString())
	})

	server.On("unsubscribe", func(c *gosocketio.Channel, ws datastructures.Workspace) {
		c.Leave(ws.ToString())
	})

	server.On("change", func(c *gosocketio.Channel, body datastructures.ChangeBody) {
		var userdata datastructures.UserData
		var err error

		if userdata, err = VerifyToken(body.Token); err != nil {
			log.Println(err)
			c.Close()
			return
		}

		if !VerifyAuthorization(body.Update.Workspace, userdata) {
			log.Println("Unauthorized")
			c.Close()
			return
		}

		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Println(err)
			return
		}
		defer ch.Close()

		if err = ch.ExchangeDeclare(
			"changes",
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		); err != nil {
			log.Println(err)
			return
		}

		var b []byte

		if b, err = json.Marshal(body.Update); err != nil {
			log.Println(err)
			return
		}

		if err = ch.Publish(
			"changes",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        b,
			}); err != nil {
			log.Println(err)
			return
		}

		c.BroadcastTo(body.Update.Workspace.ToString(), "change",
			struct {
				Sender string                    `json:"sender"`
				Update datastructures.UpdateBody `json:"update"`
			}{
				Sender: body.Sender,
				Update: body.Update,
			},
		)
	})

	return server
}
