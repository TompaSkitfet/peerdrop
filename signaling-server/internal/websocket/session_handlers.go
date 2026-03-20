package websocket

import (
	"encoding/json"
	"log"
	"math/rand"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
)

func GenerateSessionId() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVXYZ0123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (c *Client) handleCreateSession() {

	if c.peer != nil {
		c.conn.WriteJSON(Message{
			Type: "error",
			Data: "already_in_session",
		})
		return
	}

	sessionId := GenerateSessionId()

	peer := &sessions.Peer{
		Id: GenerateSessionId(),
	}

	c.sessions.Create(sessionId, peer)

	log.Println("session created:", sessionId)

	resp := Message{
		Type: "session_created",
		Data: map[string]string{
			"sessionId": sessionId,
		},
	}
	c.peer = peer
	c.conn.WriteJSON(resp)
}

func (c *Client) handleJoinSession(msg Message) {

	if c.peer != nil {
		c.conn.WriteJSON(Message{
			Type: "error",
			Data: "already_in_session",
		})
		return
	}

	var data JoinSessionData

	b, err := json.Marshal(msg.Data)
	if err != nil {
		c.conn.WriteJSON(Message{Type: "error", Data: "invalid_payload"})
		return
	}
	if err := json.Unmarshal(b, &data); err != nil {
		c.conn.WriteJSON(Message{Type: "error", Data: "invalid_payload"})
		return
	}

	peer := &sessions.Peer{
		Id: GenerateSessionId(),
	}

	session, ok := c.sessions.AddPeer(data.SessionId, peer)

	if !ok {
		c.conn.WriteJSON(Message{
			Type: "error",
			Data: "session_not_found",
		})
		return
	}

	c.peer = peer

	for p := range session.Peers {
		log.Println(p)
	}

	c.conn.WriteJSON(Message{
		Type: "session_joined",
		Data: map[string]any{
			"sessionId": session.Id,
			"peers":     len(session.Peers),
		},
	})

}
