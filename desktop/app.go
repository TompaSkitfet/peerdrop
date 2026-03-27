package main

import (
	"context"

	"github.com/TompaSkitfet/peerdrop/realtime"
	"github.com/TompaSkitfet/peerdrop/rtc"
)

// App struct
type App struct {
	ctx           context.Context
	webrtcManager *rtc.WebRTCManger
	wsManager     *realtime.WebsockerManager
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
	a.wsManager.ConnectWebSocket()
}
