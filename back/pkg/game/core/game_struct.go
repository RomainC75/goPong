package game


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
	Direction int `json:"direction"`
}

type GameStateInfos struct{
	Bait Position `json:"bait"`
	Players []Player `json:"players"`
	Level uint `json:"level"`
	GameConfig GameConfig `json:"game_config"`
}