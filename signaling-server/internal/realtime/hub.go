package realtime

type Hub struct {
	sessions      []*Session
	clients       map[*Client]bool
	register      chan *Client
	unregister    chan *Client
	createSession chan *Session
	joinSession   chan *JoinRequest
	closeSession  chan *Session
}

func NewHub() *Hub {
	return &Hub{register: make(chan *Client), unregister: make(chan *Client), createSession: make(chan *Session), joinSession: make(chan *JoinRequest), closeSession: make(chan *Session)}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
		}
	}
}
