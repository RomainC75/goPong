package managers

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
)

type GameList map[*Game]bool

type GameInterface interface {
	NewGame(manager *Manager, c *Client, r *Room )*Game
}

type GameBasicInfos struct{
	Id uuid.UUID
	Name string
}

type Game struct{
	Id uuid.UUID
	Name string
	Manager *Manager 
	Clients ClientList
	GameCore int
}

func NewGame(manager *Manager, c *Client, name string )*Game{
	newClientList := ClientList{}
	newClientList[c]=true

	return &Game{
		Id: uuid.New(),
		Name: name,
		Manager: manager,
		Clients: newClientList,
	}
}

func (g *Game) BroadcastMessage(wsMessage SocketMessage.WebSocketMessage){
	for client := range g.Clients{
		fmt.Println("send to game clients .....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}

func (g *Game) AddClient(client *Client){
	g.Clients[client]=true
}

func (g *Game) RemoveClient(client *Client){
	delete(g.Clients, client)
}

// func (gm *GameManager)CreateGame(room *Room, manager *Manager.Manager){

// 	game := NewGameInstance(c)

// 	gm.games[]
// }

