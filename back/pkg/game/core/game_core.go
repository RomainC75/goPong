package game

import (
	"fmt"
	"time"
)

func NewPlayer(playerNumber int, size int) Player{
	position := Position{0,1}
	if playerNumber > 0 {
		position.Y=-1
	}
	
	
	xShift := 4
	return Player{
		Score: 0,
		Positions: getPlayerFirstPositions(playerNumber, size, xShift, 3),
		Direction: initDirection(playerNumber),
	}
}

func NewGameState(gameStateOut chan GameStateInfos, commandArray []chan int) *GameCore{
	boardSize :=30
	gc := GameCore{
		CommandsIn: commandArray,
		GameStateOut: gameStateOut,
		GameStateInfos: GameStateInfos{
			Level: 1,
			Bait: Position{0,1},
			Players: []Player{
				NewPlayer(0, boardSize), 
				NewPlayer(1, boardSize),
			},
			GameConfig: GameConfig{
				Size: 30,
				SpeedMs: 1000,
			},
			LastCommands: []int{0,0},
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

func (gc *GameCore)MoveSnakes(){
	gc.Lock()
	for playerNumber, player := range gc.GameStateInfos.Players{
		fmt.Println(" ==================================================> player : ", playerNumber)
		// fmt.Println("player positions : ", player.Positions)
		for elIndex:= len(player.Positions)-1; elIndex>=0 ; elIndex--{
			if elIndex > 0 {
				gc.GameStateInfos.Players[playerNumber].Positions[elIndex] = gc.GameStateInfos.Players[playerNumber].Positions[elIndex-1]
			}else if elIndex == 0 {
				gc.GameStateInfos.Players[playerNumber].Direction = (gc.GameStateInfos.Players[playerNumber].Direction + gc.GameStateInfos.LastCommands[playerNumber]) % 4
				// fmt.Println("direction : ", gc.GameStateInfos.Players[playerNumber].Direction)
				if gc.GameStateInfos.Players[playerNumber].Direction == -1 {
					gc.GameStateInfos.Players[playerNumber].Direction = 3
				}
				gc.GameStateInfos.LastCommands[playerNumber] = 0
				
				switch gc.GameStateInfos.Players[playerNumber].Direction {
				case 0:
					gc.GameStateInfos.Players[playerNumber].Positions[0].X ++
				case 1:
					gc.GameStateInfos.Players[playerNumber].Positions[0].Y --
				case 2:
					gc.GameStateInfos.Players[playerNumber].Positions[0].X --
				case 3:
					gc.GameStateInfos.Players[playerNumber].Positions[0].Y ++
				}
				// fmt.Println("new head positions : ", gc.GameStateInfos.Players[playerNumber].Positions[elIndex])

			}
		}
		fmt.Println("new player positions : ", gc.GameStateInfos.Players[playerNumber].Positions[0])
	}
	gc.Unlock()
}

func (gc *GameCore)IsCollision() []bool{
	// concatenatedPositions := []Position{}
	// for _, player := range gc.GameStateInfos.Players{
	// 	concatenatedPositions = append(concatenatedPositions, player.Positions...)
	// }
	// for i, position := range concatenatedPositions{
	// 	for j, comparedPosition := range concatenatedPositions{
	// 		if i!=j && position.X == comparedPosition.X && position.Y == comparedPosition.Y{
	// 			fmt.Println("collision : ==> ", i, position, j, comparedPosition)

	// 			return true
	// 		}
	// 	}
	// }

	// ==> is Collision !!!

	res := []bool{
		false,
		false,
	}

	for _, pNumber := range []int{0,1}{
		headPosition := gc.GameStateInfos.Players[pNumber].Positions[0]
		otherPlayer := 1
		if pNumber == 1{
			otherPlayer = 0
		}
		for _, p := range gc.GameStateInfos.Players[otherPlayer].Positions{
			if headPosition.X == p.X && headPosition.Y == p.Y{
				res[pNumber] = true
			}
		}
	}
	return res

}

func (gc *GameCore)IsOutOfBoard() ([]bool){
	res := []bool{
		false,
		false,
	}
	for n, p := range gc.GameStateInfos.Players{
		size := gc.GameStateInfos.GameConfig.Size
		if p.Positions[0].X >= int(size) || p.Positions[0].X < 0 || p.Positions[0].Y >= int(size) || p.Positions[0].Y < 0 {
				res[n]=true
			}
	}
	return res
}


func (gc *GameCore)Reset(){
	xShift := 4
	for i := range gc.GameStateInfos.Players{
		gc.GameStateInfos.Players[i].Positions = getPlayerFirstPositions(i, int(gc.GameStateInfos.GameConfig.Size), xShift, 3)
		gc.GameStateInfos.Players[i].Direction = initDirection(i)
	}
	time.Sleep(time.Millisecond * time.Duration(gc.GameStateInfos.GameConfig.SpeedMs))
}


func (gc *GameCore)LaunchGameCore(){
	go func (){
		for{
			// get the client.LastValue from the core ?
			
			// set new Bait if necessary

			// Move players
			gc.MoveSnakes()
			gc.GameStateOut <- gc.GameStateInfos

			if gc.handleMistakesInGame(){
				gc.Reset()
			}
			
			time.Sleep(time.Millisecond * time.Duration(gc.GameStateInfos.GameConfig.SpeedMs))
			
		}
		// defer conn.Close()
	}()
}

// =========================================== HELPERS ===========================================


func (gc *GameCore)handleMistakesInGame() bool{
	isOut := gc.IsOutOfBoard()
	isCollision := gc.IsCollision()
	mistake:= []bool{
		false,
		false,
	}
	for i := range []int{0,1}{
		mistake[i]= isOut[i] || isCollision[i]
	}

	if mistake[0] && mistake[1]{
		fmt.Println("EQUAL MISTAKE")
		return true
	}else if mistake[0]{
		gc.GameStateInfos.Players[1].Score++
		return true
	}else if mistake[1]{
		gc.GameStateInfos.Players[0].Score++
		return true
	}
	return false
}

func initDirection(playerNumber int) int{
	if playerNumber==0{
		return 1
	}else{
		return 3
	}
}


func getPlayerFirstPositions(playerNumber int, size int, xShift int,length int) (positions []Position){
	direction := 1
	if playerNumber==0{
		direction = -1
	}
	positions = []Position{}
	for i:=0 ; i<length ; i++{
		positions = append(positions, Position{
			(size/2)-xShift+playerNumber*(xShift*2),
			(size/2) + -direction * i,
		})
	}
	return 
}