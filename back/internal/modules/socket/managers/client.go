package managers

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
)

type ClientList map[*Client]bool

type Client struct {
	id uuid.UUID
	connection *websocket.Conn
	manager *Manager
	// !!! avoid concurrent writes on the socket connection (unbuffered chan :-) )!!!
	egress chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client{
	return &Client{
		connection: conn,
		manager: manager,
		egress: make(chan []byte),
	}
}

func (c *Client) readMessages(){
	fmt.Printf("listening")
	defer func(){
		// cleanup connection
		c.manager.RemoveClient(c)
	}()
	for{
		var message SocketMessage.WebSocketMessage
		err := c.connection.ReadJSON(&message)
		if err != nil {
			log.Println("=> err : ", err.Error())
		}
		// fmt.Printf("=> ", message)
		// myChan <- message.Content["message"]

		c.manager.BroadcastMessage(message)

		// messageType, payload, err := c.connection.ReadMessage()
		// if err != nil{
		// 	if websocket.IsUnexpectedCloseError(err , websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
		// 		log.Printf("error reading message: %v", err)
		// 	}
		// 	break
		// }
		
		c.connection.WriteJSON(message)
		// // broadcast
		// for wsclient := range c.manager.clients{
		// 	wsclient.egress <- payload
		// }

		// log.Println(messageType)
		// log.Println(string(payload))
	}
}

func (c *Client) writeMessages(){
	defer func(){
		c.manager.RemoveClient(c)
	}()

	for{
		select{
		case message, ok := <- c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil{
					log.Println("connection closed:", err)
				}
				break
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil{
				log.Println("failed to send message: %v", err)
			}
			log.Println("message sent")
		}	
	}
}
