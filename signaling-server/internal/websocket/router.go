package websocket

import "log"

func (c *Client) routeMessage(msg Message) {

	switch msg.Type {

	case "create_session":
		c.handleCreateSession()

	case "join_session":
		c.handleJoinSession(msg)

	default:
		log.Println("unknown message type:", msg.Type)
	}

}
