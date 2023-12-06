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
	Game *Game
}

func NewClient(conn *websocket.Conn, manager *Manager, userData UserData) *Client{
	// TODO -> send notifications about actual state of rooms/games
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
			newRoom := c.manager.CreateRoom(message, c)
			
			// notify everyone
			bcMessage := map[string]string{
				"name": message.Content["roomName"],
				"id": (*newRoom).Id.String(),
			}
			c.manager.BroadcastMessage("ROOM_CREATED", bcMessage)

			// notify client
			backMessage:= SocketMessage.WebSocketMessage{
				Type: "ROOM_CREATED_BYYOU",
				Content: bcMessage,
			}
			c.ResponseToClient(backMessage)
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
			c.ResponseToClient(wsMessage)
		case "DISCONNECT_FROM_ROOM":
			c.manager.DisconnectUserFromRoom(c)
		case "CREATE_GAME":
			gameName := message.Content["gameName"]
			newGameId := c.manager.CreateGame(c, gameName)
			content := map[string]string{
				"id": newGameId.String(),
				"name": gameName,
			}
			wsMessage := SocketMessage.WebSocketMessage{
				Type: "GAME_CREATED_BYYOU",
				Content: content,
			}
			c.ResponseToClient(wsMessage)
			c.manager.BroadcastMessage("GAME_CREATED", content)
		case "SELECT_GAME":
			gameId, _ := uuid.Parse(message.Content["gameId"])
			c.manager.AddClientToGame(gameId, c)
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


func (c *Client)ResponseToClient(message SocketMessage.WebSocketMessage){
	m, _:= json.Marshal(message)
	c.egress <- m
}