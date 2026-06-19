package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func boolToFloat32(b bool) float32 {
	if b {
		return 1
	} else {
		return 0
	}
}

func (p *Player) checkBounds() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	if p.Center.X-p.Radius < 0 {
		p.Center.X = p.Radius
	} else if p.Center.X+p.Radius > screenWidth {
		p.Center.X = screenWidth - p.Radius
	}

	if p.Center.Y-p.Radius < 0 {
		p.Center.Y = p.Radius
	} else if p.Center.Y+p.Radius > screenHeight {
		p.Center.Y = screenHeight - p.Radius
	}
}

func (p *Player) Update(dt float32) {
	if p.InvulnerableTimer > 0 {
		p.InvulnerableTimer -= dt
	}

	p.direction.X = boolToFloat32(rl.IsKeyDown(rl.KeyD)) - boolToFloat32(rl.IsKeyDown(rl.KeyA))
	p.direction.Y = boolToFloat32(rl.IsKeyDown(rl.KeyS)) - boolToFloat32(rl.IsKeyDown(rl.KeyW))

	p.Center = rl.Vector2Add(p.Center, rl.Vector2Scale(rl.Vector2Normalize(p.direction), p.speed*dt))

	mousePos := rl.GetMousePosition()
	p.Angle = float32(math.Atan2(float64(mousePos.Y-p.Center.Y), float64(mousePos.X-p.Center.X)))

	p.checkBounds()
}
