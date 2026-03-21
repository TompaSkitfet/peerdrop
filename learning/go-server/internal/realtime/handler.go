package realtime

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	upgrader websocket.Upgrader
}

func NewHandler() *Handler {
	return &Handler{upgrader: websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}

	log.Printf("client connected")
	client := NewClient(conn)

	go client.ReadLoop()
}
