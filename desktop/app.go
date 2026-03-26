package main

import (
	"context"
	"fmt"

	"github.com/TompaSkitfet/peerdrop/rtc"
	"github.com/gorilla/websocket"
)

// App struct
type App struct {
	ctx     context.Context
	manager *rtc.WebRTCManger
	ws      *websocket.Conn
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.manager = rtc.NewManager()
	a.connectWebSocket()
}

func (a *App) connectWebSocket() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}

	a.ws = conn
	go a.listenWS()
}

func (a *App) listenWS() {
	for {
		_, msg, err := a.ws.ReadMessage()
		if err != nil {
			fmt.Println("WS error: ", err)
			return
		}
		a.handleSignal(msg)
	}
}

func (a *App) handleSignal(msg []byte) {

}
