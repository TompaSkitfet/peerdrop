package server

import (
	"github.com/TompaSkitfet/peerdrop/signaling-server/internal/websocket"
	ws "github.com/gorilla/websocket"
)

func HandleWebSocket(conn *ws.Conn) {

	client := &Client{
		Id:   websocket.GenerateSessionId(),
		conn: conn,
	}

}
