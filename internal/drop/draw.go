package drop

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (d *Drop) Draw() {
	// Blink before despawn
	if d.Life < blinkTime && int(d.Life*8)%2 == 0 {
		return
	}

	x := int32(d.Position.X)
	y := int32(d.Position.Y)

	switch d.Kind {
	case KindMedkit:
		rec := rl.NewRectangle(d.Position.X-Radius, d.Position.Y-Radius, Radius*2, Radius*2)
		rl.DrawRectangleRounded(rec, 0.3, 4, rl.RayWhite)
		rl.DrawRectangle(x-2, y-7, 4, 14, rl.Red)
		rl.DrawRectangle(x-7, y-2, 14, 4, rl.Red)
	case KindMaxHPBuff:
		rl.DrawCircleV(d.Position, Radius, color.RGBA{R: 46, G: 204, B: 113, A: 255})
		drawLabel("+10", x, y)
	case KindKnockback:
		rl.DrawCircleV(d.Position, Radius, color.RGBA{R: 230, G: 126, B: 34, A: 255})
		drawLabel("KB", x, y)
	}
}

func drawLabel(text string, x, y int32) {
	var fontSize int32 = 10
	w := rl.MeasureText(text, fontSize)
	rl.DrawText(text, x-w/2, y-fontSize/2, fontSize, rl.Black)
}
