package particle

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Particle) Update(dt float32) {
	p.Position = rl.Vector2Add(p.Position, rl.Vector2Scale(p.Velocity, dt))
	p.Velocity = rl.Vector2Scale(p.Velocity, float32(math.Pow(particleDrag, float64(dt))))
	p.Life -= dt
}
