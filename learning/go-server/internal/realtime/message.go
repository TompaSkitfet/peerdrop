package realtime

type Message struct {
	Sender      string
	MessageType string
	Content     []byte
}
