package websocket

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
	"github.com/gorilla/websocket"
)

type Handler struct {
	sessions *sessions.Manager
	upgrader websocket.Upgrader
}

func NewHandler(manager *sessions.Manager) *Handler {
	return &Handler{
		sessions: manager,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade failet: %v", err)
		return
	}

	log.Printf("client connected")
	go h.readLoop(conn)
}

func (h *Handler) readLoop(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("connection closed")
			return
		}

		var m Message

		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Println("invalid message:", err)
			continue
		}
		h.handleMessage(conn, m)
	}
}

func (h *Handler) handleMessage(conn *websocket.Conn, msg Message) {
	switch msg.Type {
	case "create_session":
		h.handleCreateSession(conn)
	case "join_session":
		h.handleJoinSession(conn, msg)
	default:
		log.Println("unknown message type: ", msg.Type)
	}
}

func generateSessionId() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVXYZ0123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (h *Handler) handleCreateSession(conn *websocket.Conn) {
	sessionId := generateSessionId()

	peer := &sessions.Peer{
		Id: generateSessionId(),
	}

	h.sessions.Create(sessionId, peer)

	log.Println("session created:", sessionId)

	resp := Message{
		Type: "session_created",
		Data: map[string]string{
			"sessonId": sessionId,
		},
	}
	conn.WriteJSON(resp)
}

func (h *Handler) handleJoinSession(conn *websocket.Conn, msg Message) {
	var data JoinSessionData

	b, _ := json.Marshal(msg.Data)
	json.Unmarshal(b, &data)

	peer := &sessions.Peer{
		Id: generateSessionId(),
	}

	session, ok := h.sessions.AddPeer(data.SessionId, peer)

	if !ok {
		conn.WriteJSON(Message{
			Type: "error",
			Data: "session_not_found",
		})
		return
	}

	conn.WriteJSON(Message{
		Type: "session_joined",
		Data: map[string]any{
			"sessionId": session.Id,
			"peers":     len(session.Peers),
		},
	})

}
