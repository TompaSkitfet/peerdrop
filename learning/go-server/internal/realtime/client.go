package realtime

import (
	"fmt"
	"math/rand"

	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	hub  *Hub
	conn *websocket.Conn
	send chan Message
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{id: generateId(), hub: hub, conn: conn, send: make(chan Message)}
}

func generateId() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (c *Client) ReadLoop() {
	defer c.conn.Close()

	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		message := Message{SenderId: c.id, MessageType: "chat", Content: m}
		c.hub.broadcast <- message
	}
}

func (c *Client) WriteLoop() {
	for message := range c.send {
		result := fmt.Sprintf("%v: %v", message.SenderId, string(message.Content))
		c.conn.WriteMessage(websocket.TextMessage, []byte(result))
	}
}
