package server

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	id   uint32
	conn *websocket.Conn
	ch   chan *[]byte
	done chan bool
}

func NewClient(id uint32, conn *websocket.Conn) *Client {
	return &Client{
		id:   id,
		conn: conn,
		ch:   make(chan *[]byte),
		done: make(chan bool),
	}
}

func (c *Client) Listen() {
	doneReading := make(chan bool)
	doneWriting := make(chan bool)

	go c.listenReadFromClient(doneReading)
	go c.listenWriteToClient(doneWriting)

	for {
		select {
		case <-c.done:
			close(doneReading)
			close(doneWriting)
			return
		}
	}
}

func (c *Client) listenWriteToClient(done chan bool) {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}()

	log.Println("Listening write to client")
	for {
		select {
		case bytes := <-c.ch:
			err := c.conn.WriteMessage(websocket.BinaryMessage, *bytes)

			if err != nil {
				log.Println(err)
			}

		case <-done:
			log.Println("Closing write goroutine")
			return
		}
	}
}

func (c *Client) listenReadFromClient(done chan bool) {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}()

	log.Println("Listening read from client")
	for {
		select {

		case <-done:
			log.Println("Closing read goroutine")
			return

		default:
			c.readFromWebSocket()
		}
	}
}

func (c *Client) readFromWebSocket() {
	messageType, message, err := c.conn.ReadMessage()
	if err != nil {
		log.Println(err)

		c.done <- true
	} else if messageType != websocket.BinaryMessage {
		log.Println("Non binary message recived, ignoring")
	} else {
		c.unmarshalUserInput(message)
	}
}

func (c *Client) unmarshalUserInput(message []byte) {
	log.Println("Recieved meesage", message)
}
