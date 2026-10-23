package game

import (
	"math/rand"

	"github.com/eswar-7116/justlive/internal/bullet"
	"github.com/eswar-7116/justlive/internal/drop"
	"github.com/eswar-7116/justlive/internal/particle"
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
			g.checkAndSaveHighScore()
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
			knockback := bulletKnockback
			if g.player.KnockbackBoost > 0 {
				knockback *= knockbackBoostMult
			}
			g.bullets = append(g.bullets, bullet.NewBullet(gunPos, dir, knockback))
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
					z.Hit(b.Damage, rl.Vector2Scale(pushDir, b.Knockback))
					g.particles = append(g.particles, particle.Spray(b.Position, b.Direction)...)
					if wasAlive && z.Health <= 0 {
						g.particles = append(g.particles, particle.Burst(z.Center())...)
						if rand.Float32() < dropChance {
							g.drops = append(g.drops, drop.RandomDrop(z.Center()))
						}
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

		for _, p := range g.particles {
			p.Update(dt)
		}

		activeParticles := g.particles[:0]
		for _, p := range g.particles {
			if p.Life > 0 {
				activeParticles = append(activeParticles, p)
			}
		}
		g.particles = activeParticles

		for _, d := range g.drops {
			d.Update(dt)
			if d.Active && rl.CheckCollisionCircles(g.player.Center, g.player.Radius, d.Position, drop.Radius) {
				if d.Apply(g.player) {
					d.Active = false
				}
			}
		}

		activeDrops := g.drops[:0]
		for _, d := range g.drops {
			if d.Active {
				activeDrops = append(activeDrops, d)
			}
		}
		g.drops = activeDrops

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
