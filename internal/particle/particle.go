package particle

import (
	"image/color"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	bloodCount   = 20
	minSpeed     = 80  // px/s
	maxSpeed     = 260 // px/s
	minLife      = 0.4 // seconds
	maxLife      = 0.8 // seconds
	minRadius    = 2   // px
	maxRadius    = 5   // px
	particleDrag = 0.02
)

type Particle struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Life     float32
	MaxLife  float32
	Radius   float32
	Color    color.RGBA
}

func randRange(lo, hi float32) float32 {
	return lo + rand.Float32()*(hi-lo)
}

func Burst(pos rl.Vector2) []*Particle {
	particles := make([]*Particle, bloodCount)
	for i := range particles {
		angle := rand.Float64() * 2 * math.Pi
		speed := randRange(minSpeed, maxSpeed)
		life := randRange(minLife, maxLife)
		particles[i] = &Particle{
			Position: pos,
			Velocity: rl.Vector2{
				X: float32(math.Cos(angle)) * speed,
				Y: float32(math.Sin(angle)) * speed,
			},
			Life:    life,
			MaxLife: life,
			Radius:  randRange(minRadius, maxRadius),
			Color:   color.RGBA{R: uint8(randRange(120, 200)), G: 0, B: 0, A: 255},
		}
	}
	return particles
}
