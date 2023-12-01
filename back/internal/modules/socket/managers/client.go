package managers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
)

type ClientList map[*Client]bool

type UserData struct {
	UserId uuid.UUID
	UserEmail string
}

type Client struct {
	userData UserData
	connection *websocket.Conn
	manager *Manager
	// !!! avoid concurrent writes on the socket connection (unbuffered chan :-) ) !!!
	egress chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager, userData UserData) *Client{
	return &Client{
		userData: userData,
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
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil{
			if websocket.IsUnexpectedCloseError(err , websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error reading message: %v", err)
			}
			break
		}
		fmt.Println(messageType, payload)


		var message SocketMessage.WebSocketMessage
		if err := json.Unmarshal(payload, &message); err != nil {
			fmt.Printf("panic reading message ! ")
        	panic(err)
    	}

		
		fmt.Printf("=> inside Client", message)
		fmt.Println("type : ", message.Type)
		// myChan <- message.Content["message"]
		
		switch message.Type {
		case "BROADCAST":
			newContent := message.Content
			newContent["userId"] = c.userData.UserId.String()
			newContent["userEmail"] = c.userData.UserEmail


			// c.manager.BroadcastMessage("BROADCAST", newContent)
			wsMessage:= SocketMessage.WebSocketMessage{
				Type: "BROADCAST",
				Content: newContent,
			}
			c.manager.broadcastC <- wsMessage

			break
		case "CREATE_ROOM":
			fmt.Println("===> CREAT_ROOM")
			backMessage := c.manager.CreateRoom(message, c)
			m, _ := json.Marshal(backMessage)
			c.egress <- m
		}
		


		
		
		// b, _ := json.Marshal(message)
		// c.egress <- b
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
