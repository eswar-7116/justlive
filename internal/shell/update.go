package shell

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Shell) Update(dt float32) {
	s.Life -= dt
	s.Position = rl.Vector2Add(s.Position, rl.Vector2Scale(s.Velocity, dt))

	if s.Grounded {
		s.Velocity = rl.Vector2Scale(s.Velocity, float32(math.Pow(groundDrag, float64(dt))))
		return
	}

	s.Rotation += s.Spin * dt
	s.VZ -= gravity * dt
	s.Z += s.VZ * dt

	// Hit the ground
	if s.Z <= 0 {
		s.Z = 0
		if -s.VZ < minBounce {
			s.Grounded = true
			s.VZ = 0
		} else {
			s.VZ = -s.VZ * bounce
			s.Spin *= 0.5
			s.Velocity = rl.Vector2Scale(s.Velocity, 0.6)
		}
	}
}
