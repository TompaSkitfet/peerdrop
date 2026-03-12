package websocket

import (
	"log"
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

		log.Printf("recieved message: %s", msg)
	}
}
