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
	egress chan []byte
	Room *Room
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
		c.manager.RemoveClient(c)
	}()
	for{
		_, payload, err := c.connection.ReadMessage()
		if err != nil{
			if websocket.IsUnexpectedCloseError(err , websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error reading message: %v", err)
			}
			break
		}
		var message SocketMessage.WebSocketMessage
		if err := json.Unmarshal(payload, &message); err != nil {
			fmt.Printf("panic reading message ! ")
        	panic(err)
    	}

		fmt.Println("=> inside Client", message)
		
		switch message.Type {
		case "BROADCAST":
			newContent := message.Content
			newContent["userId"] = c.userData.UserId.String()
			newContent["userEmail"] = c.userData.UserEmail
			c.manager.BroadcastMessage("BROADCAST", newContent)
		case "CREATE_ROOM":
			fmt.Println("===> CREAT_ROOM")
			backMessage := c.manager.CreateRoom(message, c)
			m, _ := json.Marshal(backMessage)
			c.egress <- m
		case "SEND_TO_ROOM":
			fmt.Println("send to room")
			if c.Room != nil{
				newContent := message.Content
				newContent["userId"] = c.userData.UserId.String()
				newContent["userEmail"] = c.userData.UserEmail
				wsMessage:= SocketMessage.WebSocketMessage{
					Type: "ROOM_MESSAGE",
					Content: newContent,
				}
				SendMessageToRoom(c.Room, wsMessage)
			}
		case "CONNECT_TO_ROOM":
			fmt.Println("=> try to connect to Room : ", message.Content["roomId"])
			roomUuid, _ := uuid.Parse(message.Content["roomId"])
			err := c.manager.AddUserToRoom( roomUuid, c)
			if err != nil{
				fmt.Printf("room not found ! ")
			}

			wsMessage:= SocketMessage.WebSocketMessage{
				Type: "CONNECTED_TO_ROOM",
				Content: map[string]string{"id": message.Content["roomId"],"name": c.Room.Name},
			}
			m, _ := json.Marshal(wsMessage)
			c.egress <- m
		case "DISCONNECT_FROM_ROOM":
			c.manager.DisconnectUserFromRoom(c)
		}
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
