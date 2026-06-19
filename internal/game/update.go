package game

import rl "github.com/gen2brain/raylib-go/raylib"

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

		for _, z := range g.zombies {
			z.Update(dt, g.player)
		}

		g.resolveZombieOverlaps()
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
