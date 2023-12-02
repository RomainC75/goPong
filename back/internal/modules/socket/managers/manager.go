package managers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/gorilla/websocket"
	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
	configu "github.com/saegus/test-technique-romain-chenard/pkg/configu"
)

type ManagerInterface interface{
	ServeWS(w gin.ResponseWriter , r *http.Request, userData UserData)
	AddClient(client *Client)
	RemoveClient(client *Client)
	BroadcastMessage(mType string, content map[string]string )
	CreateRoom(message SocketMessage.WebSocketMessage, client *Client)SocketMessage.WebSocketMessage
	AddUserToRoom(roomUuid uuid.UUID, client *Client) error
	DisconnectUserFromRoom(c *Client)
}

type Hub struct{
	uuid uuid.UUID
	createdAt time.Time
}

var(
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			cfg := configu.Get()
			frontUrl := cfg.Front.Host
			return origin == frontUrl
		},
	}
)

type Manager struct {
	clients ClientList
	hubs []Hub
	sync.RWMutex
	rooms RoomList
}

func New() *Manager{
	manager:= Manager{
		clients: make(ClientList),
		rooms: make(RoomList),
	}
	return &manager
}

func (m *Manager) ServeWS(w gin.ResponseWriter , r *http.Request, userData UserData){
	log.Println("new Connection")
	conn, err := websocketUpgrader.Upgrade(w,r, nil)
	if err != nil{
		log.Println(err)
		return
	}
	
	// add to client list
	client := NewClient(conn, m, userData)

	m.AddClient(client)
	
	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) AddClient(client *Client){
	// do not modify at the same time, when 2 people are trying to connect at the same time.
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) RemoveClient(client *Client){
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok{
		client.connection.Close()
		delete(m.clients, client)
	}
}


func (m *Manager) CreateRoom(message SocketMessage.WebSocketMessage, client *Client)SocketMessage.WebSocketMessage{
	m.Lock()
	defer m.Unlock()

	//create room
	roomName := message.Content["roomName"]
	newRoom := NewRoom(roomName, m, client)
	// add it to the room 
	m.rooms[newRoom]=true
	// add the room to the client
	client.Room = newRoom

	bcMessage := map[string]string{
		"name": roomName,
		"id": (*newRoom).Id.String(),
	}
	
	fmt.Println("CreateRoom : ", bcMessage)
	m.BroadcastMessage("ROOM_CREATED", bcMessage)
	
	//notify user
	backMessage:= SocketMessage.WebSocketMessage{
		Type: "ROOM_CREATED_BYYOU",
		Content: bcMessage,
	}
	return backMessage

}

func (m *Manager) AddUserToRoom(roomUuid uuid.UUID, client *Client) error {
	for room := range m.rooms{
		if room.Id.String()==roomUuid.String(){
			room.AddClient(client)
			client.Room=room

			wsMessage:= SocketMessage.WebSocketMessage{
				Type: "NEW_CONNECTION_TO_ROOM",
				Content: map[string]string{"roomId": roomUuid.String(), "userEmail": client.userData.UserEmail},
			}
			room.BroadcastMessage(wsMessage)
			return nil
		}
	}
	return errors.New("not found")
	// TODO : no room found
}

func (m *Manager) DisconnectUserFromRoom(c *Client){
	// m.Lock()
	// defer m.Unlock()

	wsMessage := SocketMessage.WebSocketMessage{
		Type: "DISCONNECTED_FROM_ROOM",
		Content: map[string]string{
			"name": c.Room.Name,
			"id": c.Room.Id.String(),
		},
	}

	room := c.Room
	// remove client from room in the roomList OR delete Room
	if len(room.Clients)==1{
		// remove room from the roomList in Manager
		delete(m.rooms, room)
	}else{
		// remove client from room in the roomList
		delete(room.Clients, c)
	}
	
	// remove room from client
	c.Room = nil
	
	// notify client
	b, _ := json.Marshal(wsMessage)
	c.egress <- b

	// broadcast in room
	wsMessage = SocketMessage.WebSocketMessage{
		Type: "USER_DISCONNECTED_FROM_ROOM",
		Content: map[string]string{
			"email": c.userData.UserEmail,
			"id": c.userData.UserId.String(),
		},
	}
	SendMessageToRoom(room, wsMessage)
}

func SendMessageToRoom(room *Room, wsMessage SocketMessage.WebSocketMessage){
	for client := range room.Clients{
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}

func (m *Manager) BroadcastMessage(mType string, content map[string]string ){
	fmt.Println("broadcast beginning function")
	// m.Lock()
	// defer m.Unlock()
	
	wsMessage := SocketMessage.WebSocketMessage{
		Type: mType,
		Content: content,
	}
	fmt.Println("==> broadcast out : ", wsMessage)

	for client := range m.clients{
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}
