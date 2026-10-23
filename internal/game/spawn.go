package game

import (
	"math/rand"

	"github.com/eswar-7116/justlive/internal/zombie"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) spawnZombie() {
	var pos rl.Vector2
	side := rand.Intn(4)
	switch side {
	case 0: // top
		pos = rl.Vector2{
			X: float32(rand.Int31n(g.width)),
			Y: -zombieSize,
		}
	case 1: // right
		pos = rl.Vector2{
			X: float32(g.width) + zombieSize,
			Y: float32(rand.Int31n(g.height)),
		}
	case 2: // bottom
		pos = rl.Vector2{
			X: float32(rand.Int31n(g.width)),
			Y: float32(g.height) + zombieSize,
		}
	case 3: // left
		pos = rl.Vector2{
			X: -zombieSize,
			Y: float32(rand.Int31n(g.height)),
		}
	}

	size := rl.Vector2{X: zombieSize, Y: zombieSize}
	z := zombie.NewZombie(pos, size, g.zombieSpeed, zombieHealth)
	g.zombies = append(g.zombies, z)
}
