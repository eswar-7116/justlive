package player

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	baseMaxHealth    float32 = 100
	baseSpeed        float32 = 250
	gunScale         float32 = 0.25
	gunDistFactor    float32 = 1.1
	ejectPortForward float32 = 0.05 // of gun length, ahead of its center
	ejectMargin      float32 = 4    // px past the body edge
)

// Death
const (
	deathFallTime  float32 = 0.35 // seconds to hit the ground
	deathSquash    float32 = 0.15 // body bulge on impact
	deathDarken    float32 = 0.45
	poolSpreadTime float32 = 1.5 // seconds
	poolSize       float32 = 1.5 // of body radius
	gunSlide       float32 = 26  // px the gun slides away
	gunSpin        float32 = 2.2 // radians the gun turns while sliding
	limbSpread     float32 = 1.1 // radians hands fall open to
)

var poolColor = color.RGBA{R: 110, G: 0, B: 0, A: 200}

// Dash
const (
	dashSpeed    float32 = 900  // px/s
	dashTime     float32 = 0.15 // seconds
	dashCooldown float32 = 1.5  // seconds
	ghostLife    float32 = 0.2  // seconds
)

type ghost struct {
	pos  rl.Vector2
	life float32
}

type Player struct {
	Center            rl.Vector2
	Radius            float32
	Color             color.RGBA
	direction         rl.Vector2
	speed             float32
	Health            float32
	MaxHealth         float32
	InvulnerableTimer float32
	Angle             float32
	KnockbackBoost    float32 // seconds left
	DashCooldown      float32 // seconds left
	dashTimer         float32
	dashDir           rl.Vector2
	ghosts            []ghost
	DeathTime         float32 // seconds since dying
}

func NewPlayer(center rl.Vector2, radius float32) *Player {
	return &Player{
		Center:            center,
		Radius:            radius,
		Color:             rl.Beige,
		speed:             baseSpeed,
		Health:            baseMaxHealth,
		MaxHealth:         baseMaxHealth,
		InvulnerableTimer: 0,
		Angle:             0,
	}
}

func (p *Player) IsDashing() bool {
	return p.dashTimer > 0
}

func (p *Player) GetGunCenter() rl.Vector2 {
	gunDist := p.Radius * gunDistFactor
	return rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(p.Angle)))*gunDist,
		Y: p.Center.Y + float32(math.Sin(float64(p.Angle)))*gunDist,
	}
}

// Ejection port on the receiver, where shells come out
func (p *Player) GetEjectionPort(gunTexture rl.Texture2D) rl.Vector2 {
	dir := rl.Vector2{
		X: float32(math.Cos(float64(p.Angle))),
		Y: float32(math.Sin(float64(p.Angle))),
	}
	gunLen := float32(gunTexture.Height) * gunScale
	dist := p.Radius*gunDistFactor + gunLen*ejectPortForward
	// Body is drawn over the back of the gun, so keep the port visible
	dist = max(dist, p.Radius+ejectMargin)
	return rl.Vector2Add(p.Center, rl.Vector2Scale(dir, dist))
}

func (p *Player) GetGunPosAndDir(gunTexture rl.Texture2D) (rl.Vector2, rl.Vector2) {
	gunPos := p.GetGunCenter()

	dir := rl.Vector2{
		X: float32(math.Cos(float64(p.Angle))),
		Y: float32(math.Sin(float64(p.Angle))),
	}

	gunLen := float32(gunTexture.Height) * gunScale
	tip := rl.Vector2Add(gunPos, rl.Vector2Scale(dir, gunLen/2))

	return tip, dir
}
