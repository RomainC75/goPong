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

func NewGameState(gameStateOut chan GameStateInfos, commandArray []chan int) *GameCore{
	gc := GameCore{
		CommandsIn: commandArray,
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
			LastCommands: []int{10,10},
		},
		
	}
	gc.LaunchGameCore()
	gc.LaunchCommandListener(0)
	gc.LaunchCommandListener(1)
	return &gc
}

func (gc *GameCore)LaunchCommandListener(PlayerNumber int){
	go func(){
		for{
			select{
			case command  := <- gc.CommandsIn[PlayerNumber]:
				fmt.Println("=> WRITE COMMAND DATA !!", command)
				gc.Lock()
				gc.GameStateInfos.LastCommands[PlayerNumber] = command
				gc.Unlock()
			}
		}
	}()
}


func (gc *GameCore)LaunchGameCore(){
	go func (){
		for{
			// get the client.LastValue from the core ?
			
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