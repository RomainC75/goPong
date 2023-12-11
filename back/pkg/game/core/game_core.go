package game

import (
	"fmt"
	"time"
)

type GameCore struct {
	CommandIn chan CommandMessage
	GameStateOut chan GameStateInfos
	GameStateInfos GameStateInfos
	GameConfig GameConfig
}

type GameConfig struct{
	size uint
	speedMs uint 
}

type Position struct{
	X float64 `json:"x"`
	Y float64 `jons:"y"`
}

type CommandMessage struct{
	PlayerNumber int
	Command string
}

type Player struct {
	Score int `json:"score"`
	Positions []Position `json:"positions"`
}

type GameStateInfos struct{
	Bait Position `json:"bait"`
	Players []Player `json:"players"`
	Level uint `json:"level"`
}

func NewPlayer(number int) Player{
	position := Position{0,1}
	if number > 0 {
		position.Y=-1
	}
	return Player{
		Score: 0,
		Positions: []Position{
			Position{0,1},
		},
	}
}

func NewGameState(commandIn chan CommandMessage, gameStateOut chan GameStateInfos) *GameCore{
	gc := GameCore{
		CommandIn: commandIn,
		GameStateOut: gameStateOut,
		GameStateInfos: GameStateInfos{
			Level: 1,
			Bait: Position{0,1},
			Players: []Player{
				NewPlayer(0), 
				NewPlayer(1),
			},
		},
		GameConfig: GameConfig{
			size: 30,
			speedMs: 1000,
		},
	}
	gc.LaunchGameCore()
	return &gc
}

func (gc *GameCore)LaunchGameCore(){
	fmt.Printf("GAME CORE CREATION")
	go func (){
		for{
			fmt.Println("GAME CORE LOOP")
			select{
			case messageIn, _ := <- gc.CommandIn:
				fmt.Println("messageIn : ", messageIn, gc.GameStateInfos.Bait.X)
			default:
			}

			gc.GameStateInfos.Bait.X += 1
			// fmt.Println("message received : ", message, ball.Position.Bait.X)
			gc.GameStateOut <- gc.GameStateInfos

			time.Sleep(time.Millisecond * time.Duration(gc.GameConfig.speedMs))

			// var message SocketMessage.Message
			// err := conn.ReadJSON(&message)
			// if !errors.Is(err, nil) {
			// 	log.Printf("error occurred: %v", err)
			// 	break
			// }

			// fmt.Println("MESSAGEZ : ", message.Message)

			// conn.WriteMessage(websocket.TextMessage, []byte(message.Message))
		}
		// defer conn.Close()
	}()
}