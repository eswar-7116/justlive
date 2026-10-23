package particle

import (
	"image/color"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const particleDrag = 0.02

// Death burst
const (
	bloodCount = 20
	minSpeed   = 80  // px/s
	maxSpeed   = 260 // px/s
	minLife    = 0.4 // seconds
	maxLife    = 0.8 // seconds
	minRadius  = 2   // px
	maxRadius  = 5   // px
)

// Hit spray
const (
	sprayCount     = 6
	sprayBackCount = 2                  // of sprayCount, toward the shooter
	sprayAngle     = 35 * math.Pi / 180 // radians each side of bullet direction
	sprayMinSpeed  = 150                // px/s
	sprayMaxSpeed  = 350                // px/s
	sprayMinLife   = 0.2                // seconds
	sprayMaxLife   = 0.4                // seconds
	sprayMinRadius = 1.5                // px
	sprayMaxRadius = 3                  // px
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

func newBlood(pos rl.Vector2, angle float64, speed, life, radius float32) *Particle {
	return &Particle{
		Position: pos,
		Velocity: rl.Vector2{
			X: float32(math.Cos(angle)) * speed,
			Y: float32(math.Sin(angle)) * speed,
		},
		Life:    life,
		MaxLife: life,
		Radius:  radius,
		Color:   color.RGBA{R: uint8(randRange(120, 200)), G: 0, B: 0, A: 255},
	}
}

func Burst(pos rl.Vector2) []*Particle {
	particles := make([]*Particle, bloodCount)
	for i := range particles {
		angle := rand.Float64() * 2 * math.Pi
		particles[i] = newBlood(pos, angle,
			randRange(minSpeed, maxSpeed),
			randRange(minLife, maxLife),
			randRange(minRadius, maxRadius),
		)
	}
	return particles
}

func Spray(pos, dir rl.Vector2) []*Particle {
	baseAngle := math.Atan2(float64(dir.Y), float64(dir.X))
	particles := make([]*Particle, sprayCount)
	for i := range particles {
		angle := baseAngle + (rand.Float64()*2-1)*sprayAngle
		if i < sprayBackCount {
			angle += math.Pi
		}
		particles[i] = newBlood(pos, angle,
			randRange(sprayMinSpeed, sprayMaxSpeed),
			randRange(sprayMinLife, sprayMaxLife),
			randRange(sprayMinRadius, sprayMaxRadius),
		)
	}
	return particles
}
