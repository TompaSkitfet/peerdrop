package sessions

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Peer struct {
	Id   string
	Conn *websocket.Conn
}

type Session struct {
	Id    string
	Peers map[string]*Peer
}

type Manager struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

func NewManager() *Manager {
	return &Manager{sessions: make(map[string]*Session)}
}

func (m *Manager) Create(sessionId string, peer *Peer) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()

	s := &Session{
		Id:    sessionId,
		Peers: map[string]*Peer{peer.Id: peer},
	}

	m.sessions[sessionId] = s
	return s
}

func (m *Manager) Get(sessionId string) (*Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	s, ok := m.sessions[sessionId]
	return s, ok
}

func (m *Manager) AddPeer(sessionId string, peer *Peer) (*Session, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s, ok := m.sessions[sessionId]
	if !ok {
		return nil, false
	}
	s.Peers[peer.Id] = peer
	return s, true
}

func (m *Manager) RemovePeer(sessionId string, peerId string) {

	m.mu.Lock()
	defer m.mu.Unlock()

	s, ok := m.sessions[sessionId]
	if !ok {
		return
	}

	delete(s.Peers, peerId)

	if len(s.Peers) == 0 {
		delete(m.sessions, sessionId)
	}
}
