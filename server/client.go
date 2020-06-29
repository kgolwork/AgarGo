package server

import "github.com/gorilla/websocket"

type Client struct {
	id   uint32
	conn *websocket.Conn
}

func NewClient(id uint32, conn *websocket.Conn) *Client {
	return &Client{
		id:   id,
		conn: conn,
	}
}

func (c *Client) ListenWrite() {

}
