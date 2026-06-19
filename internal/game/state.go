package game

type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
	GameStateOver
)
