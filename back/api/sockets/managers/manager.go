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
	SocketMessage "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	config "github.com/saegus/test-technique-romain-chenard/config"
	GameCore "github.com/saegus/test-technique-romain-chenard/pkg/game/core"
)

type ManagerInterface interface{
	ServeWS(w gin.ResponseWriter , r *http.Request, userData UserData)
	AddClient(client *Client)
	RemoveClient(client *Client)
	BroadcastMessage(mType string, content map[string]string )
	CreateRoom(message SocketMessage.WebSocketMessage, client *Client) *Room
	AddUserToRoom(roomUuid uuid.UUID, client *Client) error
	DisconnectUserFromRoom(c *Client)

	CreateGame(c *Client, name string) uuid.UUID
	AddClientToGame(gameId uuid.UUID,c *Client)error
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
			cfg := config.Get()
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
	games GameList
}

func New() *Manager{
	manager:= Manager{
		clients: make(ClientList),
		rooms: make(RoomList),
		games: make(GameList),
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
	m.NotifyClientStateOfRoomsAndGames(client)
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


func (m *Manager) CreateRoom(message SocketMessage.WebSocketMessage, client *Client) *Room{
	m.Lock()
	defer m.Unlock()

	//create room
	roomName := message.Content["roomName"]
	newRoom := NewRoom(roomName, m, client)
	// add it to the room 
	m.rooms[newRoom]=true
	// add the room to the client
	client.Room = newRoom

	return newRoom
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

func (m *Manager) BroadcastMessage(mType string, content map[string]string ){
	// m.Lock()
	// defer m.Unlock()
	
	wsMessage := SocketMessage.WebSocketMessage{
		Type: mType,
		Content: content,
	}
	
	for client := range m.clients{
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}

func (m *Manager) CreateGame(c *Client, name string) uuid.UUID{
	// create Room in manager for game
	// message := SocketMessage.WebSocketMessage{
	// 	Content: map[string]string{
	// 		"name": c.userData.UserEmail,
	// 	},
	// }
	// room := m.CreateRoom(message, c)
	
	// add game to manager
	newGame := NewGame(m, c, name)
	m.games[newGame]=true
	
	// add game to client
	c.Game=newGame
	
	return newGame.Id
}
	
func (m *Manager) AddClientToGame(gameId uuid.UUID,c *Client)error{
	for game := range m.games{
		if game.Id.String()==gameId.String() && game.Full==false {
			// add game to client
			c.Game=game
			
			// add client to game
			game.AddClient(c)
			fmt.Println("CLIENT added ! :", game.Clients)

			if len(game.Clients)==game.MaxPlayerNumber{
				game.Full=true
				fmt.Printf("====================> LAUNCH GAME !! ")
				game.GameCore = GameCore.NewGameState(game.CommandIn, game.GameStateOut)
			}
			
			return nil
		}
	}
	return errors.New("game not found")
}
		
func SendMessageToRoom(room *Room, wsMessage SocketMessage.WebSocketMessage){
	for client := range room.Clients{
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}

func (m *Manager) GetActualRoomsBasicInfos() []RoomBasicInfos{
	rooms := []RoomBasicInfos{}
	for room := range m.rooms{
		rooms = append(rooms, RoomBasicInfos{
			Id: room.Id,
			Name: room.Name,
		})
	}
	return rooms
}

func (m *Manager) GetActualGamesBasicInfos() []GameBasicInfos{
	games := []GameBasicInfos{}
	for game := range m.games{
		games = append(games, GameBasicInfos{
			Id: game.Id,
			Name: game.Name,
		})
	}
	return games
}

func (m *Manager) NotifyClientStateOfRoomsAndGames(c *Client){
	rooms := m.GetActualRoomsBasicInfos()
	bRooms, _ := json.Marshal(rooms)
	games := m.GetActualGamesBasicInfos()
	bGames, _ := json.Marshal(games)
	wsMessage := SocketMessage.WebSocketMessage{
		Type: "ROOMS_GAMES_NOTIFICATION",
		Content: map[string]string{
			"rooms": string(bRooms),
			"games": string(bGames),
		},
	}
	b, _ := json.Marshal(wsMessage)
	c.egress <- b
}