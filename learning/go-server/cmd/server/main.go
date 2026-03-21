package main

import (
	"log"
	"net/http"

	"github.com/TompaSkitfet/peerdrop/go-server/internal/realtime"
)

func main() {
	hub := realtime.NewHub()
	go hub.Run()

	handler := realtime.NewHandler(hub)

	mux := http.NewServeMux()
	mux.Handle("/ws", handler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
