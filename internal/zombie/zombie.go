package zombie

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Zombie struct {
	Position  rl.Vector2
	Size      rl.Vector2
	Color     color.RGBA
	direction rl.Vector2
	speed     float32
	health    float32
	knockback rl.Vector2
}

func NewZombie(position rl.Vector2, size rl.Vector2, speed float32, health float32) *Zombie {
	return &Zombie{
		Position: position,
		Size:     size,
		Color:    rl.Red,
		speed:    speed,
		health:   health,
	}
}
