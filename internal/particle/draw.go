package particle

import rl "github.com/gen2brain/raylib-go/raylib"

func (p *Particle) Draw() {
	c := p.Color
	c.A = uint8(255 * max(p.Life/p.MaxLife, 0))
	rl.DrawCircleV(p.Position, p.Radius, c)
}
