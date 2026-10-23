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
	if p.KnockbackBoost > 0 {
		p.KnockbackBoost -= dt
	}

	if p.DashCooldown > 0 {
		p.DashCooldown -= dt
	}

	p.direction.X = boolToFloat32(rl.IsKeyDown(rl.KeyD)) - boolToFloat32(rl.IsKeyDown(rl.KeyA))
	p.direction.Y = boolToFloat32(rl.IsKeyDown(rl.KeyS)) - boolToFloat32(rl.IsKeyDown(rl.KeyW))

	mousePos := rl.GetMousePosition()
	p.Angle = float32(math.Atan2(float64(mousePos.Y-p.Center.Y), float64(mousePos.X-p.Center.X)))

	if rl.IsKeyPressed(rl.KeySpace) && p.DashCooldown <= 0 && !p.IsDashing() {
		p.startDash()
	}

	if p.IsDashing() {
		p.dashTimer -= dt
		p.Center = rl.Vector2Add(p.Center, rl.Vector2Scale(p.dashDir, dashSpeed*dt))
		p.ghosts = append(p.ghosts, ghost{pos: p.Center, life: ghostLife})
	} else {
		p.Center = rl.Vector2Add(p.Center, rl.Vector2Scale(rl.Vector2Normalize(p.direction), p.speed*dt))
	}

	p.updateGhosts(dt)
	p.checkBounds()
}

func (p *Player) startDash() {
	// Dash toward aim when standing still
	if p.direction.X == 0 && p.direction.Y == 0 {
		p.dashDir = rl.Vector2{X: float32(math.Cos(float64(p.Angle))), Y: float32(math.Sin(float64(p.Angle)))}
	} else {
		p.dashDir = rl.Vector2Normalize(p.direction)
	}
	p.dashTimer = dashTime
	p.DashCooldown = dashCooldown
}

func (p *Player) updateGhosts(dt float32) {
	activeGhosts := p.ghosts[:0]
	for _, g := range p.ghosts {
		g.life -= dt
		if g.life > 0 {
			activeGhosts = append(activeGhosts, g)
		}
	}
	p.ghosts = activeGhosts
}
