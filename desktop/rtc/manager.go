package rtc

import (
	"sync"

	"github.com/pion/webrtc/v3"
)

type WebRTCManger struct {
	peer *Peer
	mu   sync.RWMutex
	emit func(string, interface{})
}

func NewManager() *WebRTCManger {
	return &WebRTCManger{}
}

func (m *WebRTCManger) CreatePeer(id string) error {
	cfg := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{{
			URLs: []string{"stun:stun.l.google.com:19302"},
		}},
	}

	pc, err := webrtc.NewPeerConnection(cfg)
	if err != nil {
		return err
	}

	pc.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		m.emit("ice-candidate", c.ToJSON())
	})

	pc.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		m.emit("connection-state", s.String())
	})

	m.peer = &Peer{
		PC: pc,
	}

	return nil
}

func (m *WebRTCManger) SendFile(path string) error
func (m *WebRTCManger) ClosePeer()
