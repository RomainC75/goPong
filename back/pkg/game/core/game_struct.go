package game

import (
	"sync"
)

type GameCore struct {
	sync.RWMutex
	CommandsIn []chan int
	GameStateOut chan GameStateInfos
	GameStateInfos GameStateInfos
}

type GameConfig struct{
	Size uint `json:"size"`
	SpeedMs uint `json:"speed_ms"`
}

type Position struct{
	X int `json:"x"`
	Y int `json:"y"`
}

type CommandMessage struct{
	PlayerNumber int
	Command int
}

type Player struct {
	Score int `json:"score"`
	Positions []Position `json:"positions"`
	Direction int `json:"direction"`
}

type GameStateInfos struct{
	Bait Position `json:"bait"`
	Players []Player `json:"players"`
	Level uint `json:"level"`
	GameConfig GameConfig `json:"game_config"`
	LastCommands []int `json:"last_command"`
}