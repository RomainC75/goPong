package managers

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client{
	return &Client{
		connection: conn,
		manager: manager,
	}
}

func (c *Client) readMessages(){
	fmt.Printf("listening")
	defer func(){
		// cleanup connection
		c.manager.RemoveClient(c)
	}()
	for{
		var message SocketMessage.Message
		err := c.connection.ReadJSON(&message)
		if err != nil {
			log.Println("=> err : ", err.Error())
		}
		fmt.Printf("=> ", message)
		// messageType, payload, err := c.connection.ReadMessage()

		// if err != nil{
		// 	if websocket.IsUnexpectedCloseError(err , websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
		// 		log.Printf("error reading message: %v", err)
		// 	}
		// 	break
		// }

		// log.Println(messageType)
		// log.Println(string(payload))
	}
}