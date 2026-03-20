package websocket

import "log"

func (c *Client) routeMessage(msg Message) {

	switch msg.Type {

	case "create_session":
		c.handleCreateSession()

	case "join_session":
		c.handleJoinSession(msg)

	case "leave_session":
		c.handleLeaveSession()

	case "signal":
		c.handleSignal()

	default:
		log.Println("unknown message type:", msg.Type)
	}

}
