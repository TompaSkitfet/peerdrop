package main

import (
	"log"
	"net/http"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/websocket"
)

func main() {

	sessionsManager := sessions.NewManager()
	wsHandler := websocket.NewHandler(sessionsManager)

	mux := http.NewServeMux()
	mux.Handle("/ws", wsHandler)

	addr := ":8080"

	log.Printf("signaling server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}

}
