package realtime

import "github.com/TompaSkitfet/peerdrop/go-server/internal/util"

type Session struct {
	id         string
	title      string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
}

func NewSession(title string) *Session {
	return &Session{
		id:         util.GenerateId(),
		title:      title,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
}

func (s *Session) Run() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			delete(s.clients, client)
		case message := <-s.broadcast:
			for client := range s.clients {
				client.send <- message
			}
		}
	}
}
