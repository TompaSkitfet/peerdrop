package realtime

import (
	"fmt"
	"strings"

	"github.com/TompaSkitfet/peerdrop/go-server/internal/util"
	"github.com/gorilla/websocket"
)

type Client struct {
	id       string
	userName string
	hub      *Hub
	session  *Session
	conn     *websocket.Conn
	send     chan Message
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	id := util.GenerateId()
	return &Client{id: id, userName: id, hub: hub, conn: conn, send: make(chan Message)}
}

func (c *Client) ReadLoop() {
	defer c.conn.Close()

	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.parseMessage(string(m))
	}
}

func (c *Client) WriteLoop() {
	for message := range c.send {
		switch message.MessageType {
		case "system":
			c.conn.WriteMessage(websocket.TextMessage, []byte(message.Content))

		case "chat":
			result := fmt.Sprintf("%v: %v", message.Sender, string(message.Content))
			c.conn.WriteMessage(websocket.TextMessage, []byte(result))
		}
	}
}

func (c *Client) parseMessage(m string) {
	if strings.HasPrefix(m, "/") {
		c.parseSystemMessage(m)
		return
	}
	if c.session == nil {
		return
	}
	message := Message{Sender: c.userName, MessageType: "chat", Content: []byte(m)}
	c.session.broadcast <- message
}

func (c *Client) parseSystemMessage(m string) {
	parts := strings.SplitN(m, " ", 2)
	if len(parts) < 2 {
		c.conn.WriteMessage(websocket.TextMessage, []byte("Invalid arguments"))
	}

	cmd := parts[0]
	arg := strings.TrimSpace(parts[1])

	switch cmd {
	case "/name":
		c.hub.systemBroadcast <- c.setUsername(arg)
	case "/create":
		request := CreateSessionRequest{c, arg}
		c.hub.create <- request
	case "/join":
		request := JoinSessionRequest{c, arg}
		c.hub.join <- request
	}
}

func (c *Client) setUsername(name string) Message {
	currentName := c.userName
	newName := name

	c.userName = newName

	m := fmt.Sprintf("%s changed name to %s", currentName, newName)
	return Message{Sender: c.userName, MessageType: "system", Content: []byte(m)}
}
