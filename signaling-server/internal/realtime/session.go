package realtime

type Session struct {
	Id   string
	Host *Client
	Peer *Client
}

type JoinSessionRequest struct {
	client    *Client
	sessionId string
}
