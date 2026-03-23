package realtime

type Session struct {
	Id   string
	Host *Client
	Peer *Client
}

type JoinRequest struct {
	client    *Client
	isHost    bool
	sessionId string
}
