package utils

import (
	"log"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"

	datastructures "../data_structures"
)

// NewChangeServer creates a socket server for async file change transactions with the client
func NewChangeServer() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("Connected")
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})

	server.On("subscribe", func(c *gosocketio.Channel, ws datastructures.Workspace) {
		token := c.RequestHeader()["Authentication"][0]

		var userdata datastructures.UserData
		var err error

		if userdata, err = VerifyToken(token); err != nil {
			log.Println(err)
			c.Close()
			return
		}

		if !VerifyAuthorization(ws, userdata) {
			log.Println("Unauthorized")
			c.Close()
			return
		}

		c.Join(ws.ToString())
		log.Println("Sub")
	})

	server.On("unsubscribe", func(c *gosocketio.Channel, ws datastructures.Workspace) {
		c.Leave(ws.ToString())
		log.Println("Unsub")
	})

	server.On("change", func(c *gosocketio.Channel, change datastructures.UpdateBody) {
		c.BroadcastTo(change.Workspace.ToString(), "change", change)
	})

	return server
}
