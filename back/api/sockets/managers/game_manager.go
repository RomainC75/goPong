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
	p1CommandIn chan GameCore.CommandMessage
	p2CommandIn chan GameCore.CommandMessage
	GameStateOut chan GameCore.GameStateInfos
}

func NewGame(manager *Manager, c *Client, name string )*Game{
	newClientList := []*Client{c}
	c.PlayerNumber = 0

	
	g := &Game{
		Id: uuid.New(),
		Name: name,
		Manager: manager,
		Clients: newClientList,
		MaxPlayerNumber: 2,
		// p1CommandIn: make(chan GameCore.CommandMessage, 1000),
		// p2CommandIn: make(chan GameCore.CommandMessage, 1000),
		GameStateOut: make(chan GameCore.GameStateInfos, 1000),
	}
	// c.CommandIn = g.p1CommandIn
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
	// ---- creation
	commandArray := []chan int{
		make( chan int ),
		make( chan int ),
	}

	client.PlayerNumber = 1
	
	g.Clients = append(g.Clients, client)
	clientIds := []UserData{}
	for i, client := range g.Clients{
		client.CommandIn = commandArray[i]
		clientIds = append(clientIds, UserData{
			UserId: client.userData.UserId,
			UserEmail: client.userData.UserEmail,
		})
	}
	// allocation to gameCore
	g.Full=true
	g.GameCore = GameCore.NewGameState(g.GameStateOut, commandArray)



	bConfig, _ := json.Marshal(g.GameCore.GameStateInfos.GameConfig)
	bClients, _ := json.Marshal(clientIds)

	message := SocketMessage.WebSocketMessage{
		Type: "GAME_CONFIG_BROADCAST",
		Content: map[string]string{
			"clients": string(bClients),
			"config": string(bConfig),
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

func (g *Game) writeMessages(){
	// defer func(){
	// 	c.manager.RemoveClient(c)
	// }()

	for{
		select{
		case message, _ := <- g.GameStateOut:
			// utils.PrettyDisplay(message)
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

