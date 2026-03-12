package websocket

import (
	"encoding/json"
	"log"
	"math/rand"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
)

func generateSessionId() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVXYZ0123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (c *Client) handleCreateSession() {
	sessionId := generateSessionId()

	peer := &sessions.Peer{
		Id: generateSessionId(),
	}

	c.sessions.Create(sessionId, peer)

	log.Println("session created:", sessionId)

	resp := Message{
		Type: "session_created",
		Data: map[string]string{
			"sessionId": sessionId,
		},
	}
	c.conn.WriteJSON(resp)
}

func (c *Client) handleJoinSession(msg Message) {
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
		Id: generateSessionId(),
	}

	session, ok := c.sessions.AddPeer(data.SessionId, peer)

	if !ok {
		c.conn.WriteJSON(Message{
			Type: "error",
			Data: "session_not_found",
		})
		return
	}

	c.conn.WriteJSON(Message{
		Type: "session_joined",
		Data: map[string]any{
			"sessionId": session.Id,
			"peers":     len(session.Peers),
		},
	})

}
