package game

import (
	"github.com/eswar-7116/justlive/internal/bullet"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) update(dt float32) {
	switch g.gameState {
	case GameStateMenu:
		if rl.IsKeyDown(rl.KeyP) {
			g.gameState = GameStatePlaying
		}
	case GameStatePlaying:
		if g.player.Health <= 0 {
			g.gameState = GameStateOver
			break
		}

		g.player.Update(dt)

		g.spawnTimer += dt
		if g.spawnTimer >= g.spawnInterval {
			g.spawnZombie()
			g.spawnTimer = 0
		}

		if g.shootTimer > 0 {
			g.shootTimer -= dt
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) && g.shootTimer <= 0 {
			gunPos, dir := g.player.GetGunPosAndDir(g.gunTexture)
			g.bullets = append(g.bullets, bullet.NewBullet(gunPos, dir))
			g.shootTimer = g.fireRate
		}

		for _, b := range g.bullets {
			if b.Active {
				b.Update(dt)
			}
		}

		for _, b := range g.bullets {
			if !b.Active {
				continue
			}
			for _, z := range g.zombies {
				if rl.CheckCollisionCircleRec(b.Position, b.Radius, z.CollisionRec()) {
					wasAlive := z.Health > 0
					z.Health -= b.Damage
					if wasAlive && z.Health <= 0 {
						g.score += 10
						if g.score > 0 && g.score%50 == 0 {
							g.zombieSpeed += 10
							for _, zo := range g.zombies {
								zo.IncreaseSpeed(10)
							}
						}
					}
					b.Active = false

					pushDir := rl.Vector2Normalize(b.Direction)
					z.Knockback = rl.Vector2Add(z.Knockback, rl.Vector2Scale(pushDir, 500))
					break
				}
			}
		}

		for _, z := range g.zombies {
			z.Update(dt, g.player)
		}

		g.resolveZombieOverlaps()

		activeBullets := g.bullets[:0]
		for _, b := range g.bullets {
			if b.Active {
				activeBullets = append(activeBullets, b)
			}
		}
		g.bullets = activeBullets

		activeZombies := g.zombies[:0]
		for _, z := range g.zombies {
			if z.Health > 0 {
				activeZombies = append(activeZombies, z)
			}
		}
		g.zombies = activeZombies
	}
}

func (g *Game) resolveZombieOverlaps() {
	for i := 0; i < len(g.zombies); i++ {
		for j := i + 1; j < len(g.zombies); j++ {
			a := g.zombies[i]
			b := g.zombies[j]

			centerA := rl.Vector2{
				X: a.Position.X + a.Size.X/2,
				Y: a.Position.Y + a.Size.Y/2,
			}

			centerB := rl.Vector2{
				X: b.Position.X + b.Size.X/2,
				Y: b.Position.Y + b.Size.Y/2,
			}

			diff := rl.Vector2Subtract(centerB, centerA)
			dist := rl.Vector2Length(diff)

			minDist := (a.Size.X + b.Size.X) / 2

			if dist > 0 && dist < minDist {
				push := ((minDist - dist) / 2) + 2

				dir := rl.Vector2Normalize(diff)

				a.Position = rl.Vector2Subtract(
					a.Position,
					rl.Vector2Scale(dir, push),
				)

				b.Position = rl.Vector2Add(
					b.Position,
					rl.Vector2Scale(dir, push),
				)
			}
		}
	}
}
