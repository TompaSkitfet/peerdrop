package websocket

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data,omitempty"`
}

type JoinSessionData struct {
	SessionId string `json:"sessionId"`
}
