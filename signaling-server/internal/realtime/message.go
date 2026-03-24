package realtime

import "encoding/json"

type Message struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

type MessageType string

const (
	CreateSession MessageType = "create_session"
	JoinSession   MessageType = "join_session"
	Offer         MessageType = "offer"
	Answer        MessageType = "answer"
	IceCandidate  MessageType = "ice-candidate"
)

func (t MessageType) IsValid() bool {
	switch t {
	case CreateSession, JoinSession, Offer, Answer, IceCandidate:
		return true
	default:
		return false
	}
}

type JoinSessionData struct {
	SessionId string `json:"session_id"`
}

type SDPData struct {
	SDP string `json:"sdp"`
}

type ICECandidateDate struct {
	Candidate     string `json:"candidate"`
	SDPMid        string `json:"sdpMid"`
	SDPMLineIndex int    `json:"sdpMLineIndex"`
}
