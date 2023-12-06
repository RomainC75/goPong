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
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Game struct{
	Id uuid.UUID
	Name string
	Manager *Manager 
	Clients []*Client
	GameCore int
	MaxPlayerNumber int
	Full bool
}

func NewGame(manager *Manager, c *Client, name string )*Game{
	newClientList := []*Client{c}

	return &Game{
		Id: uuid.New(),
		Name: name,
		Manager: manager,
		Clients: newClientList,
		MaxPlayerNumber: 2,
	}
}

func (g *Game) BroadcastMessage(wsMessage SocketMessage.WebSocketMessage){
	for _, client := range g.Clients{
		fmt.Println("send to game clients .....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}

func (g *Game) AddClient(client *Client){
	g.Clients = append(g.Clients, client)
	clientIds := []UserData{}
	for _, client := range g.Clients{
		clientIds = append(clientIds, UserData{
			UserId: client.userData.UserId,
			UserEmail: client.userData.UserEmail,
		})
	}
	bClients, _ := json.Marshal(clientIds)
	message := SocketMessage.WebSocketMessage{
		Type: "GAME_BROADCAST",
		Content: map[string]string{
			"clients": string(bClients),
		},
	}
	g.BroadcastMessage(message)
}

func (g *Game) RemoveClient(client *Client){
	for i, c := range g.Clients{
		if c.userData.UserId.String()==client.userData.UserId.String(){
			g.Clients = append(g.Clients[:i], g.Clients[i+1:]...)
		}
	}
}

// func (gm *GameManager)CreateGame(room *Room, manager *Manager.Manager){

// 	game := NewGameInstance(c)

// 	gm.games[]
// }

