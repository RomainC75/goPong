package game

import (
	"fmt"
	"time"
)

type GameCore struct {
	CommandIn chan CommandMessage
	GameStateOut chan GameStateInfos
	GameStateInfos GameStateInfos
}

type GameConfig struct{
	Size uint `json:"size"`
	SpeedMs uint `json:"speed_ms"`
}

type Position struct{
	X float64 `json:"x"`
	Y float64 `json:"y"`
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
	GameConfig GameConfig `json:"game_config"`
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
			GameConfig: GameConfig{
				Size: 30,
				SpeedMs: 1000,
			},
		},
		
	}
	gc.LaunchGameCore()
	return &gc
}

func (gc *GameCore)LaunchGameCore(){
	go func (){
		for{
			select{
			case messageIn, _ := <- gc.CommandIn:
				fmt.Println("messageIn : ", messageIn, gc.GameStateInfos.Bait.X)
			default:
			}
			gc.GameStateInfos.Bait.X += 1
			gc.GameStateOut <- gc.GameStateInfos
			time.Sleep(time.Millisecond * time.Duration(gc.GameStateInfos.GameConfig.SpeedMs))
			// var message SocketMessage.Message
			// err := conn.ReadJSON(&message)
			// if !errors.Is(err, nil) {
			// 	break
			// }
			// conn.WriteMessage(websocket.TextMessage, []byte(message.Message))
		}
		// defer conn.Close()
	}()
}