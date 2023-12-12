package game

import (
	"fmt"
	"time"
)


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

func NewGameState(p1CommandIn chan CommandMessage, p2CommandIn chan CommandMessage, gameStateOut chan GameStateInfos) *GameCore{
	gc := GameCore{
		p1CommandIn: p1CommandIn,
		p2CommandIn: p2CommandIn,
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
			case messageIn, _ := <- gc.p1CommandIn:
				
				
				// for messageIn := range gc.p1CommandIn {
					
				// }
				fmt.Println("messageIn P1: ", messageIn, gc.GameStateInfos.Bait.X)
			default:
			}

			select{
			case messageIn, _ := <- gc.p2CommandIn:
				
				// for messageIn := range gc.p2CommandIn {
				// 	if messageIn.PlayerNumber
				// }
				fmt.Println("messageIn P2: ", messageIn, gc.GameStateInfos.Bait.X)
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