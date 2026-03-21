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
		Id:   GenerateSessionId(),
		Conn: c.conn,
	}

	c.sessions.Create(sessionId, peer)

	log.Println("session created:", sessionId)

	resp := Message{
		Type: "session_created",
		Data: map[string]string{
			"sessionId": sessionId,
			"peerId":    peer.Id,
		},
	}
	c.peer = peer
	c.sessionId = sessionId

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
		Id:   GenerateSessionId(),
		Conn: c.conn,
	}

	session, ok := c.sessions.AddPeer(data.SessionId, peer)

	if !ok {
		c.conn.WriteJSON(Message{
			Type: "error",
			Data: "session_not_found",
		})
		return
	}

	for id, p := range session.Peers {
		if id == c.peer.Id {
			continue
		}

		p.Conn.WriteJSON(Message{
			Type: "peer_joined",
			Data: map[string]string{
				"peerId": c.peer.Id,
			},
		})
	}

	c.peer = peer
	c.sessionId = session.Id

	peerIds := []string{}
	for id := range session.Peers {
		if id != c.peer.Id {
			peerIds = append(peerIds, id)
		}
	}

	c.conn.WriteJSON(Message{
		Type: "session_joined",
		Data: map[string]any{
			"sessionId": session.Id,
			"peerId":    peer.Id,
			"peers":     peerIds,
		},
	})

}

func (c *Client) handleLeaveSession() {
	if c.peer == nil || c.sessionId == "" {
		return
	}

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

	c.peer = nil
	c.sessionId = ""

	c.conn.WriteJSON(Message{
		Type: "session_left",
	})
}

func (c *Client) handleSignal(msg Message) {
	if c.peer == nil || c.sessionId == "" {
		return
	}

	var data SignalData
	b, err := json.Marshal(msg.Data)
	if err != nil {
		return
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return
	}

	session, ok := c.sessions.Get(c.sessionId)
	if !ok {
		return
	}

	for id, peer := range session.Peers {
		if id == c.peer.Id {
			continue
		}
		if data.To != "" && data.To != id {
			continue
		}

		peer.Conn.WriteJSON(Message{
			Type: "signal",
			Data: map[string]any{
				"from": c.peer.Id,
				"data": data.Data,
			},
		})
	}
}
