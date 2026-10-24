package shell

import rl "github.com/gen2brain/raylib-go/raylib"

func (s *Shell) Draw() {
	alpha := min(max(s.Life/fadeTime, 0), 1)
	if s.Grounded {
		alpha *= groundAlpha
	}

	// Top-down camera: height makes it bigger, not offset
	scale := 1 + s.Z/zScale
	w := width * scale
	h := height * scale
	origin := rl.Vector2{X: w / 2, Y: h / 2}

	body := brass
	body.A = uint8(255 * alpha)
	rl.DrawRectanglePro(rl.Rectangle{X: s.Position.X, Y: s.Position.Y, Width: w, Height: h}, origin, s.Rotation, body)

	// Same origin keeps the rim on the shell's back end
	end := rim
	end.A = uint8(255 * alpha)
	rl.DrawRectanglePro(rl.Rectangle{X: s.Position.X, Y: s.Position.Y, Width: rimWidth * scale, Height: h}, origin, s.Rotation, end)
}
