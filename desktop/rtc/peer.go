package rtc

import "github.com/pion/webrtc/v3"

type Peer struct {
	Id string

	PC *webrtc.PeerConnection
	DC *webrtc.DataChannel

	SessionDir string
	Transfers  map[string]*IncomingFile
}
