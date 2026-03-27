package realtime

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type WebsockerManager struct {
	Conn *websocket.Conn
}

func (m *WebsockerManager) ConnectWebSocket() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}

	m.Conn = conn
	go m.listenWS()
}

func (m *WebsockerManager) listenWS() {
	for {
		_, msg, err := m.Conn.ReadMessage()
		if err != nil {
			fmt.Println("WS error: ", err)
			return
		}
		m.handleSignal(msg)
	}
}

func (m *WebsockerManager) handleSignal(msg []byte) {

}
