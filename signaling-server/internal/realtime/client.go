package realtime

import "github.com/gorilla/websocket"

type Client struct {
	Id     string
	Hub    *Hub
	Conn   *websocket.Conn
	Sesion *Session
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{Id: "", Hub: hub, Conn: conn}
}
