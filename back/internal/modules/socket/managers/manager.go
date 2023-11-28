package managers

import (
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
	ServeWS(w gin.ResponseWriter , r *http.Request)
	AddClient(client *Client)
	RemoveClient(client *Client)
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
}

func New() *Manager{
	return &Manager{
		clients: make(ClientList),
	}
}

func (m *Manager) ServeWS(w gin.ResponseWriter , r *http.Request){
	log.Println("new Connection")
	conn, err := websocketUpgrader.Upgrade(w,r, nil)
	if err != nil{
		log.Println(err)
		return
	}
	
	// add to client list
	client := NewClient(conn, m)


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

func (m *Manager) BroadcastMessage(message SocketMessage.WebSocketMessage){
	m.Lock()
	defer m.Unlock()

	newMessage := SocketMessage.WebSocketMessage{
		Type: "BroadCast",
		Content: message.Content,
	}

	for client := range m.clients{
		client.connection.WriteJSON(newMessage)
	}
}