package managers

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
)


type RoomList map[*Room]bool

type RoomBasicInfos struct{
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}
type Room struct{
	Id uuid.UUID
	Name string
	RoomBasicInfos
	Manager *Manager
	Clients ClientList
}

func NewRoom(name string, manager *Manager, client *Client) *Room{
	clients := ClientList{}
	clients[client]=true
	return &Room{
		Id: uuid.New(),
		Name: name,
		Manager: manager,
		Clients: clients,
	}
}

func (r *Room)AddClient(client *Client){
	r.Clients[client]=true
}

func (r *Room)RemoveClient(client *Client){
	delete(r.Clients,client)
}

func (r *Room)BroadcastMessage(wsMessage  SocketMessage.WebSocketMessage){
	for client := range r.Clients{
		// client.connection.WriteJSON(newMessage)
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}
