package main

import (
	"context"

	"github.com/TompaSkitfet/peerdrop/rtc"
	"github.com/TompaSkitfet/peerdrop/signaling"
)

// App struct
type App struct {
	ctx           context.Context
	webrtcManager *rtc.WebRTCManger
	wsManager     *signaling.WebsocketManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.webrtcManager = rtc.NewManager()
	a.wsManager = signaling.NewWebsocketManager(ctx)
	a.wsManager.ConnectWebSocket()
}
