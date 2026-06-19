package bullet

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	Position  rl.Vector2
	Direction rl.Vector2
	Speed     float32
	Damage    float32
	Radius    float32
	Active    bool
}

func NewBullet(pos, dir rl.Vector2) *Bullet {
	return &Bullet{
		Position:  pos,
		Direction: rl.Vector2Normalize(dir),
		Speed:     1000,
		Damage:    25,
		Radius:    5,
		Active:    true,
	}
}

func (b *Bullet) Update(dt float32) {
	b.Position = rl.Vector2Add(b.Position, rl.Vector2Scale(b.Direction, b.Speed*dt))
	
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	
	if b.Position.X < 0 || b.Position.X > screenWidth || b.Position.Y < 0 || b.Position.Y > screenHeight {
		b.Active = false
	}
}

func (b *Bullet) Draw(tex rl.Texture2D) {
	scale := float32(0.5)
	destRec := rl.Rectangle{
		X:      b.Position.X,
		Y:      b.Position.Y,
		Width:  float32(tex.Width) * scale,
		Height: float32(tex.Height) * scale,
	}

	sourceRec := rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height))
	origin := rl.Vector2{
		X: (float32(tex.Width) * scale) / 2,
		Y: (float32(tex.Height) * scale) / 2,
	}
	
	angle := float32(math.Atan2(float64(b.Direction.Y), float64(b.Direction.X)))
	rotation := (angle * 180 / math.Pi) + 90 // Adjust rotation depending on the bullet image

	rl.DrawTexturePro(tex, sourceRec, destRec, origin, rotation, rl.White)
}
