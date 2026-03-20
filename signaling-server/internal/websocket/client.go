package websocket

import (
	"encoding/json"
	"log"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn      *websocket.Conn
	sessions  *sessions.Manager
	peer      *sessions.Peer
	sessionId string
}

func NewClient(conn *websocket.Conn, sessions *sessions.Manager) *Client {
	return &Client{
		conn:     conn,
		sessions: sessions,
	}
}

func (c *Client) ReadLoop() {
	defer func() {
		if c.peer != nil && c.sessionId != "" {

			session, ok := c.sessions.Get(c.sessionId)
			if ok {
				for id, p := range session.Peers {
					if id == c.peer.Id {
						continue
					}

					p.Conn.WriteJSON(Message{
						Type: "peer_left",
						Data: map[string]string{
							"peerId": c.peer.Id,
						},
					})
				}
			}

			c.sessions.RemovePeer(c.sessionId, c.peer.Id)
		}
		c.conn.Close()
	}()

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
