package game

import (
	"fmt"
	"time"
)

type GameCore struct {
	CommandIn chan CommandMessage
	GameStateOut chan GameStateMessage
	Ball Ball
}

type Ball struct {
	Position Position
	Direction float64
}

type Position struct{
	X float64
	Y float64
}

type CommandMessage struct{
	PlayerNumber int
	Command string
}

type GameStateMessage struct{
	Ball Ball
}

func NewGameState(commandIn chan CommandMessage, gameStateOut chan GameStateMessage) *GameCore{

	gc := GameCore{
		CommandIn: commandIn,
		GameStateOut: gameStateOut,
	}
	gc.LaunchGameCore()
	return &gc
}

func (gc *GameCore)LaunchGameCore(){
	ball := Ball{}
	fmt.Printf("GAME CORE CREATION")
	go func (){
		for{
			fmt.Printf("GAME CORE LOOP")
			select{
				
			case message, _ := <- gc.CommandIn:
				ball.Position.X += 1
				fmt.Println("message received : ", message, ball.Position.X)
				gc.GameStateOut <- GameStateMessage{
					Ball: ball,
				}	
			default:
				
			}

			time.Sleep(time.Millisecond * 1000)

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