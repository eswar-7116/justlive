package zombie

import (
	"math"

	"github.com/eswar-7116/justlive/internal/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (z *Zombie) Update(dt float32, p *player.Player) {
	zrect := rl.Rectangle{
		X:      z.Position.X,
		Y:      z.Position.Y,
		Width:  z.Size.X,
		Height: z.Size.Y,
	}

	if rl.CheckCollisionCircleRec(p.Center, p.Radius, zrect) {
		p.Health -= 20

		pushDir := rl.Vector2Normalize(
			rl.Vector2Subtract(z.Position, p.Center),
		)

		z.knockback = rl.Vector2Scale(pushDir, 300)
	}

	z.direction = rl.Vector2Subtract(p.Center, z.Position)

	moveVel := rl.Vector2Scale(
		rl.Vector2Normalize(z.direction),
		z.speed,
	)

	totalVel := rl.Vector2Add(moveVel, z.knockback)

	z.Position = rl.Vector2Add(
		z.Position,
		rl.Vector2Scale(totalVel, dt),
	)
	z.knockback = rl.Vector2Scale(z.knockback, float32(math.Pow(0.01, float64(dt))))
}
