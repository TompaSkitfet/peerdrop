package realtime

import "log"

type Hub struct {
	clients         map[*Client]bool
	sessions        map[string]*Session
	create          chan CreateSessionRequest
	join            chan JoinSessionRequest
	register        chan *Client
	unregister      chan *Client
	systemBroadcast chan Message
}

type CreateSessionRequest struct {
	client *Client
	title  string
}

type JoinSessionRequest struct {
	client *Client
	id     string
}

func NewHub() *Hub {
	return &Hub{clients: map[*Client]bool{}, sessions: map[string]*Session{}, create: make(chan CreateSessionRequest), join: make(chan JoinSessionRequest), register: make(chan *Client), unregister: make(chan *Client), systemBroadcast: make(chan Message)}
}

func (h *Hub) Run() {
	for {
		select {
		case sessionRequest := <-h.create:
			h.createSession(sessionRequest)
		case sessionRequest := <-h.join:
			h.joinSession(sessionRequest)
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
		case message := <-h.systemBroadcast:
			for client := range h.clients {
				client.send <- message
			}
		}
	}
}

func (h *Hub) createSession(sr CreateSessionRequest) {
	session := NewSession(sr.title)
	session.clients[sr.client] = true
	h.sessions[session.id] = session
	sr.client.session = session
	go session.Run()
	log.Printf("Created session: %v", session.id)
}

func (h *Hub) joinSession(sr JoinSessionRequest) {
	session := h.sessions[sr.id]
	session.clients[sr.client] = true
	sr.client.session = session
}
