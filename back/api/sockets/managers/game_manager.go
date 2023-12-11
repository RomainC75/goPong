package managers

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	SocketMessage "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	GameCore "github.com/saegus/test-technique-romain-chenard/pkg/game/core"
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
	GameCore *GameCore.GameCore
	MaxPlayerNumber int
	Full bool
	CommandIn chan GameCore.CommandMessage
	GameStateOut chan GameCore.GameStateInfos
}

func NewGame(manager *Manager, c *Client, name string )*Game{
	newClientList := []*Client{c}
	
	g := &Game{
		Id: uuid.New(),
		Name: name,
		Manager: manager,
		Clients: newClientList,
		MaxPlayerNumber: 2,
		CommandIn: make(chan GameCore.CommandMessage, 1000),
		GameStateOut: make(chan GameCore.GameStateInfos, 1000),
	}
	go g.writeMessages()
	return g
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
	g.GameCore = GameCore.NewGameState(g.CommandIn, g.GameStateOut)
}

func (g *Game) RemoveClient(client *Client){
	for i, c := range g.Clients{
		if c.userData.UserId.String()==client.userData.UserId.String(){
			g.Clients = append(g.Clients[:i], g.Clients[i+1:]...)
		}
	}
}

func (g *Game) writeMessages(){
	// defer func(){
	// 	c.manager.RemoveClient(c)
	// }()

	for{
		select{
		case message, _ := <- g.GameStateOut:

			slcState, _ := json.Marshal(message)

			g.BroadcastMessage(SocketMessage.WebSocketMessage{
				Type: "GAME_STATE",
				Content: map[string]string{
					"state": string(slcState),
				},
			})
		}	
	}
}

