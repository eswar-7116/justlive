package bot

import (
	"github.com/eswar-7116/justlive/internal/input"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game -> bot (sent every frame)
type State struct {
	State   string        `json:"state"` // menu, playing, over
	Width   int32         `json:"width"`
	Height  int32         `json:"height"`
	Score   int           `json:"score"`
	Player  PlayerState   `json:"player"`
	Zombies []ZombieState `json:"zombies"`
	Drops   []DropState   `json:"drops"`
}

type PlayerState struct {
	X              float32 `json:"x"`
	Y              float32 `json:"y"`
	Radius         float32 `json:"radius"`
	HP             float32 `json:"hp"`
	MaxHP          float32 `json:"maxHp"`
	Angle          float32 `json:"angle"`
	DashReady      bool    `json:"dashReady"`
	KnockbackBoost float32 `json:"knockbackBoost"`
}

type ZombieState struct {
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	HP float32 `json:"hp"`
}

type DropState struct {
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Kind string  `json:"kind"` // medkit, maxHpBuff, knockback
}

// Bot -> game
type Action struct {
	Move       [2]float32 `json:"move"`
	Aim        [2]float32 `json:"aim"`
	Shoot      bool       `json:"shoot"`
	DashCount  uint32     `json:"dashCount"`
	StartCount uint32     `json:"startCount"`
}

func (a Action) heldInput() input.Input {
	return input.Input{
		Move:  rl.Vector2{X: a.Move[0], Y: a.Move[1]},
		Aim:   rl.Vector2{X: a.Aim[0], Y: a.Aim[1]},
		Shoot: a.Shoot,
	}
}
