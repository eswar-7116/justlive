package player

import rl "github.com/gen2brain/raylib-go/raylib"

func (p *Player) Draw() {
	rl.DrawCircleV(p.Center, p.Radius, p.Color)
}
