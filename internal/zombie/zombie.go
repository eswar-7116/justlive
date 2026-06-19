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
	Health    float32
	Knockback rl.Vector2
	animTimer float32
	animFrame int
}

func NewZombie(position rl.Vector2, size rl.Vector2, speed float32, health float32) *Zombie {
	return &Zombie{
		Position:  position,
		Size:      size,
		Color:     rl.Red,
		speed:     speed,
		Health:    health,
		animTimer: 0,
		animFrame: 0,
	}
}

func (z *Zombie) DrawRec() rl.Rectangle {
	const scale float32 = 3.5
	width := z.Size.X * scale
	height := z.Size.Y * scale
	x := z.Position.X - (width-z.Size.X)/2
	y := z.Position.Y - (height-z.Size.Y)/2

	return rl.Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (z *Zombie) CollisionRec() rl.Rectangle {
	drawRec := z.DrawRec()
	colWidth := drawRec.Width * 0.40
	colHeight := drawRec.Height * 0.72
	colX := drawRec.X + (drawRec.Width-colWidth)/2
	colY := drawRec.Y + (drawRec.Height-colHeight)/2

	return rl.Rectangle{
		X:      colX,
		Y:      colY,
		Width:  colWidth,
		Height: colHeight,
	}
}

func (z *Zombie) IncreaseSpeed(amount float32) {
	z.speed += amount
}
