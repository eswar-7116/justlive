package player

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Center            rl.Vector2
	Radius            float32
	Color             color.RGBA
	direction         rl.Vector2
	speed             float32
	Health            float32
	InvulnerableTimer float32
	Angle             float32
}

func NewPlayer(center rl.Vector2, radius float32) *Player {
	return &Player{
		Center:            center,
		Radius:            radius,
		Color:             rl.Beige,
		speed:             250,
		Health:            100,
		InvulnerableTimer: 0,
		Angle:             0,
	}
}

func (p *Player) GetGunPosAndDir(gunTexture rl.Texture2D) (rl.Vector2, rl.Vector2) {
	gunDist := p.Radius * 1.1
	gunScale := float32(0.2)
	gunPos := rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(p.Angle)))*gunDist,
		Y: p.Center.Y + float32(math.Sin(float64(p.Angle)))*gunDist,
	}

	dir := rl.Vector2{
		X: float32(math.Cos(float64(p.Angle))),
		Y: float32(math.Sin(float64(p.Angle))),
	}

	gunLen := float32(gunTexture.Height) * gunScale
	tip := rl.Vector2Add(gunPos, rl.Vector2Scale(dir, gunLen/2))

	return tip, dir
}
