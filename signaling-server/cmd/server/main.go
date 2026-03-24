package main

import (
	"log"
	"net/http"

	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/realtime"
)

func main() {
	hub := realtime.NewHub()
	go hub.Run()

	handler := realtime.NewHandler(hub)

	mux := http.NewServeMux()
	mux.Handle("/ws", handler)

	addr := ":8080"

	log.Printf("signaling server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
