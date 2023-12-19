package game

import (
	"errors"
	"fmt"
	"math/rand"
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

func (gc *GameCore)MoveSnakes(isStretch bool, stretchPlayer int){
	if isStretch{
		fmt.Println("STRETCH !!")
	}

	gc.Lock()
	for playerNumber, player := range gc.GameStateInfos.Players{
		lastPart := Position{}
	
		if isStretch && stretchPlayer==playerNumber{
			lastPart = player.Positions[len(player.Positions)-1]
		}

		for elIndex := len(player.Positions)-1 ; elIndex>=0 ; elIndex--{
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
		if isStretch && stretchPlayer == playerNumber{
			gc.GameStateInfos.Players[playerNumber].Positions = append(gc.GameStateInfos.Players[playerNumber].Positions, lastPart)
		}
	}
	gc.Unlock()
}

func (gc *GameCore)IsCollision() []bool{
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

func (gc *GameCore)IsBaitEaten() (int, error){
	for i, p := range gc.GameStateInfos.Players{
		hp := p.Positions[0]
		bp := gc.GameStateInfos.Bait
		if hp.X == bp.X && hp.Y == bp.Y{
			return i, nil
		}
	}
	return 0, errors.New("no bait eaten")
}

func (gc *GameCore)CreateNewBait(){
	newP := Position{
		X: rand.Intn(int(gc.GameStateInfos.GameConfig.Size)),
		Y: rand.Intn(int(gc.GameStateInfos.GameConfig.Size)),
	}
	for _, p := range gc.GameStateInfos.Players{
		for _, pPart := range p.Positions{
			if pPart.X == newP.X && pPart.Y == newP.Y{
				gc.CreateNewBait()
				return 
			}
		}
	}
	gc.GameStateInfos.Bait = newP
}


func (gc *GameCore)Reset(){
	xShift := 4
	for i := range gc.GameStateInfos.Players{
		gc.GameStateInfos.Players[i].Positions = getPlayerFirstPositions(i, int(gc.GameStateInfos.GameConfig.Size), xShift, 3)
		gc.GameStateInfos.Players[i].Direction = initDirection(i)
	}
	time.Sleep(time.Millisecond * time.Duration(gc.GameStateInfos.GameConfig.SpeedMs))
}

func (gc *GameCore)ScoreUp(playerNumber int){
	gc.GameStateInfos.Players[playerNumber].Score ++
}

func (gc *GameCore)LaunchGameCore(){
	go func (){
		for{
			playerNum, err := gc.IsBaitEaten()
			if err != nil {
				gc.MoveSnakes(false, 0)
			}else{
				gc.ScoreUp(playerNum)
				gc.CreateNewBait()
				gc.MoveSnakes(true, playerNum)
			}
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