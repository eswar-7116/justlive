package game

import (
	"github.com/eswar-7116/justlive/internal/bullet"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) update(dt float32) {
	switch g.gameState {
	case GameStateMenu:
		if rl.IsKeyPressed(rl.KeyP) {
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
			g.spawnInterval = max(g.spawnInterval*spawnIntervalDecay, minSpawnInterval)
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
					pushDir := rl.Vector2Normalize(b.Direction)
					z.Hit(b.Damage, rl.Vector2Scale(pushDir, bulletKnockback))
					if wasAlive && z.Health <= 0 {
						g.score += killScore
						if g.score > 0 && g.score%speedUpScore == 0 {
							g.zombieSpeed = min(g.zombieSpeed+zombieSpeedStep, maxZombieSpeed)
							for _, zo := range g.zombies {
								zo.IncreaseSpeed(zombieSpeedStep, maxZombieSpeed)
							}
						}
					}
					b.Active = false
					break
				}
			}
		}

		hpBefore := g.player.Health
		for _, z := range g.zombies {
			z.Update(dt, g.player)
		}
		if g.player.Health < hpBefore {
			g.shake = min(g.shake+shakePerHit, 1)
		}
		g.shake = max(g.shake-shakeDecay*dt, 0)

		g.resolveZombieOverlaps(dt)

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
	case GameStateOver:
		if rl.IsKeyPressed(rl.KeyP) {
			g.reset()
			g.gameState = GameStatePlaying
		}
	}
}

func (g *Game) resolveZombieOverlaps(dt float32) {
	for i := 0; i < len(g.zombies); i++ {
		for j := i + 1; j < len(g.zombies); j++ {
			a := g.zombies[i]
			b := g.zombies[j]

			diff := rl.Vector2Subtract(b.Center(), a.Center())
			dist := rl.Vector2Length(diff)

			minDist := (a.Size.X + b.Size.X) / 2

			if dist > 0 && dist < minDist {
				push := ((minDist - dist) / 2) + zombieJostle*dt

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
