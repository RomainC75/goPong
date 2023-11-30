package managers

import (
	"encoding/json"
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
	BroadcastMessage(mType string, message map[string]string)
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
	return &Manager{
		clients: make(ClientList),
	}
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


func (m *Manager) CreateRoom(message SocketMessage.WebSocketMessage, client *Client){
	m.Lock()
	defer m.Unlock()

	//create room
	roomName := message.Content["roomName"]
	newRoom := NewRoom(roomName, m, client)
	// add it to the room 
	m.rooms[newRoom]=true

	bcMessage := map[string]string{
		"name": roomName,
		"id": (*newRoom).Id.String(),
	}
	m.BroadcastMessage("ROOM_CREATED", bcMessage)
	//notify users

}

// func (m *Manager) AddUserToRoom(){

// }

// func (m *Manager) DeleteRoom(message SocketMessage.WebSocketMessage){

// }



func (m *Manager) BroadcastMessage(mType string, message map[string]string){
	m.Lock()
	defer m.Unlock()

	newMessage := SocketMessage.WebSocketMessage{
		Type: mType,
		Content: message,
	}

	for client := range m.clients{
		// client.connection.WriteJSON(newMessage)
		b, _ := json.Marshal(newMessage)
		client.egress <- b
	}
}
