package signaling

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type WebsocketManager struct {
	Conn *websocket.Conn
	ctx  context.Context
}

func NewWebsocketManager(ctx context.Context) *WebsocketManager {
	return &WebsocketManager{ctx: ctx}
}

func (m *WebsocketManager) ConnectWebSocket() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}

	m.Conn = conn
	go m.listenWS()
}

func (m *WebsocketManager) listenWS() {
	for {
		_, msg, err := m.Conn.ReadMessage()
		if err != nil {
			fmt.Println("WS error: ", err)
			return
		}
		m.handleSignal(msg)
	}
}

func (m *WebsocketManager) handleSignal(msg []byte) {
	runtime.EventsEmit(m.ctx, "signal", string(msg))
}

func (m *WebsocketManager) SendSignal(msg string) error {
	return m.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}
