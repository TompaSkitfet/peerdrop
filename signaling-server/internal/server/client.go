package server

import (
	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/sessions"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id      string
	conn    *websocket.Conn
	session *sessions.Session
}
