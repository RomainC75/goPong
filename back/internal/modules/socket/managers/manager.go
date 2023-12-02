package managers

import (
	"encoding/json"
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
	BroadcastMessage()
	CreateRoom(message SocketMessage.WebSocketMessage, client *Client)SocketMessage.WebSocketMessage
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
	broadcastC chan SocketMessage.WebSocketMessage
}

func New() *Manager{
	manager:= Manager{
		clients: make(ClientList),
		rooms: make(RoomList),
		broadcastC: make(chan SocketMessage.WebSocketMessage),
	}
	go manager.BroadcastMessage()
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

	// m.BroadcastMessage(, bcMessage)
	wsMessage:= SocketMessage.WebSocketMessage{
		Type: "ROOM_CREATED",
		Content: bcMessage,
	}
	m.broadcastC <- wsMessage


	fmt.Println("post broadcast")
	
	//notify user
	backMessage:= SocketMessage.WebSocketMessage{
		Type: "ROOM_CREATED_BYYOU",
		Content: bcMessage,
	}
	return backMessage

}

func (m *Manager) AddUserToRoom(client *Client){

}

// func (m *Manager) DeleteRoom(message SocketMessage.WebSocketMessage){

// }

func SendMessageToRoom(room *Room, wsMessage SocketMessage.WebSocketMessage){
	for client := range room.Clients{
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}

func (m *Manager) BroadcastMessage(){
	// TODO : try withou m.broadcastC
	for{
		select{
		case wsMessage, _ := <- m.broadcastC:
			// mType string, message map[string]string
			fmt.Println("broadcast beginning function")
			// m.Lock()
			// defer m.Unlock()
			fmt.Println("lock ? ")
			
			fmt.Println("==> broadcast out : ", wsMessage)

			for client := range m.clients{
				// client.connection.WriteJSON(newMessage)
				fmt.Println("send.....")
				b, _ := json.Marshal(wsMessage)
				client.egress <- b
			}
		}
	}
	
}
