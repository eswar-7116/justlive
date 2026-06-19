package zombie

import rl "github.com/gen2brain/raylib-go/raylib"

func (z *Zombie) Draw() {
	rl.DrawRectangleV(z.Position, z.Size, z.Color)
}
