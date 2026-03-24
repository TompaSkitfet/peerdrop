package realtime

import "log"

type Hub struct {
	sessions      map[string]*Session
	clients       map[*Client]bool
	register      chan *Client
	unregister    chan *Client
	createSession chan *Session
	joinSession   chan JoinSessionRequest
	closeSession  chan *Session
}

func NewHub() *Hub {
	return &Hub{sessions: map[string]*Session{}, clients: map[*Client]bool{}, register: make(chan *Client), unregister: make(chan *Client), createSession: make(chan *Session), joinSession: make(chan JoinSessionRequest), closeSession: make(chan *Session)}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
		case session := <-h.createSession:
			h.sessions[session.Id] = session
			log.Printf("session created: %s", session.Id)
		case request := <-h.joinSession:
			if _, ok := h.sessions[request.sessionId]; !ok {
				request.client.send <- []byte("Session doesn't exist")
				continue
			}
			if h.sessions[request.sessionId].Peer != nil {
				request.client.send <- []byte("Session full")
				continue
			}
			h.sessions[request.sessionId].Peer = request.client
			request.client.Session = h.sessions[request.sessionId]
		}
	}
}
