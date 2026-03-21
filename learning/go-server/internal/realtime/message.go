package realtime

type Message struct {
	SenderId    string
	MessageType string
	Content     []byte
}
