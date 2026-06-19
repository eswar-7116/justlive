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
			Y: -30,
		}
	case 1: // right
		pos = rl.Vector2{
			X: float32(g.width) + 30,
			Y: float32(rand.Int31n(g.height)),
		}
	case 2: // bottom
		pos = rl.Vector2{
			X: float32(rand.Int31n(g.width)),
			Y: float32(g.height) + 30,
		}
	case 3: // left
		pos = rl.Vector2{
			X: -30,
			Y: float32(rand.Int31n(g.height)),
		}
	}

	size := rl.Vector2{X: 30, Y: 30}
	z := zombie.NewZombie(pos, size, 50, 100)
	g.zombies = append(g.zombies, z)
}
