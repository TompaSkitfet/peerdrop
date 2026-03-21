package realtime

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) ReadLoop() {
	defer c.conn.Close()

	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		c.conn.WriteMessage(messageType, message)
	}
}
