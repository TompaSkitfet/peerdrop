package websocket

import (
	"encoding/json"
	"log"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	sessions *sessions.Manager
}

func NewClient(conn *websocket.Conn, sessions *sessions.Manager) *Client {
	return &Client{
		conn:     conn,
		sessions: sessions,
	}
}

func (c *Client) ReadLoop() {
	defer c.conn.Close()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("connection closed")
			return
		}

		var m Message

		if err := json.Unmarshal(msg, &m); err != nil {
			log.Println("invalid message:", err)
			continue
		}

		c.routeMessage(m)
	}
}
