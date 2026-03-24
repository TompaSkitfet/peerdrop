package realtime

import (
	"encoding/json"
	"log"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/util"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id      string
	Hub     *Hub
	Conn    *websocket.Conn
	Session *Session
	send    chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	id := util.GenerateId()
	return &Client{
		Id:   id,
		Hub:  hub,
		Conn: conn,
		send: make(chan []byte, 256),
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		c.parseMessage(m)
	}
}

func (c *Client) WritePump() {
	for msg := range c.send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}

func (c *Client) parseMessage(msg []byte) {
	var m Message

	if err := json.Unmarshal(msg, &m); err != nil {
		log.Printf("invalid request")
		return
	}

	switch m.Type {
	case CreateSession:
		c.Hub.createSession <- c.createSession()
	case JoinSession:
		r, err := c.joinSession(m.Data)
		if err != nil {
			log.Printf("invalid request")
			return
		}
		c.Hub.joinSession <- r
	case SDP:
		c.parseSDP(m)
		return
	case IceCandidate:
		return
	}
}

func (c *Client) createSession() *Session {
	id := util.GenerateId()
	session := &Session{Id: id, Host: c}
	c.Session = session
	return session
}

func (c *Client) joinSession(data json.RawMessage) (JoinSessionRequest, error) {
	var r JoinSessionData
	if err := json.Unmarshal(data, &r); err != nil {
		return JoinSessionRequest{}, err
	}
	return JoinSessionRequest{sessionId: r.SessionId, client: c}, nil
}

func (c *Client) parseSDP(msg Message) {
	var reciever *Client
	if c == c.Session.Host {
		reciever = c.Session.Peer
	} else {
		reciever = c.Session.Host
	}

	if c.Session.Peer == nil {
		c.send <- []byte("no peer connected")
		return
	}
	var r SDPData
	if err := json.Unmarshal(msg.Data, &r); err != nil {
		c.send <- []byte("invalid offer request")
		return
	}

	data, err := json.Marshal(msg.Data)
	if err != nil {
		c.send <- []byte("failed to marshal offer")
		return
	}
	reciever.send <- []byte(data)
}
