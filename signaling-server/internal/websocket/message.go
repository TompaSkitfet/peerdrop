package websocket

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data,omitempty"`
}

type JoinSessionData struct {
	SessionId string `json:"sessionId"`
}

type SignalData struct {
	To   string `json:"to,omitempty"`
	Data any    `json:"data"`
}
