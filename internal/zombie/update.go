package zombie

import (
	"math"

	"github.com/eswar-7116/justlive/internal/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (z *Zombie) Update(dt float32, p *player.Player) {
	zrect := z.CollisionRec()

	if rl.CheckCollisionCircleRec(p.Center, p.Radius, zrect) {
		if p.InvulnerableTimer <= 0 {
			p.Health -= contactDamage
			p.InvulnerableTimer = contactInvuln
		}

		pushDir := rl.Vector2Normalize(
			rl.Vector2Subtract(z.Center(), p.Center),
		)

		z.Knockback = rl.Vector2Scale(pushDir, contactKnockback)
	}

	z.direction = rl.Vector2Subtract(p.Center, z.Center())

	moveVel := rl.Vector2Scale(
		rl.Vector2Normalize(z.direction),
		z.speed,
	)

	totalVel := rl.Vector2Add(moveVel, z.Knockback)

	z.Position = rl.Vector2Add(
		z.Position,
		rl.Vector2Scale(totalVel, dt),
	)
	z.Knockback = rl.Vector2Scale(z.Knockback, float32(math.Pow(0.01, float64(dt))))

	z.animTimer += dt
	if z.animTimer >= 0.1 {
		z.animFrame = (z.animFrame + 1) % 10
		z.animTimer -= 0.1
	}
}
