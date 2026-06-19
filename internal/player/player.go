package player

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Center    rl.Vector2
	Radius    float32
	Color     color.RGBA
	direction rl.Vector2
	speed     float32
	Health    float32
}

func NewPlayer(center rl.Vector2, radius float32) *Player {
	return &Player{
		Center: center,
		Radius: radius,
		Color:  rl.Beige,
		speed:  250,
		Health: 100,
	}
}
