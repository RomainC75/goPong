package managers

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager *Manager

	// avoid concurrent writes on the socket connection
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
		// var message SocketMessage.WebSocketMessage
		// err := c.connection.ReadJSON(&message)
		// if err != nil {
		// 	log.Println("=> err : ", err.Error())
		// }
		// fmt.Printf("=> ", message)
		// myChan <- message.Content["message"]

		messageType, payload, err := c.connection.ReadMessage()

		if err != nil{
			if websocket.IsUnexpectedCloseError(err , websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error reading message: %v", err)
			}
			break
		}
		
		log.Println(messageType)
		log.Println(string(payload))
	}
}

