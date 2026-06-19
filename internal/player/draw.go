package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) Draw(gunTexture rl.Texture2D) {
	color := p.Color
	if p.InvulnerableTimer > 0 {
		if int(p.InvulnerableTimer*15)%2 == 0 {
			color = rl.Red
		}
	}

	borderThickness := float32(2.5)

	// Draw Gun
	gunDist := p.Radius * 1.1
	gunPos := rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(p.Angle)))*gunDist,
		Y: p.Center.Y + float32(math.Sin(float64(p.Angle)))*gunDist,
	}

	gunScale := float32(0.2)
	destRec := rl.Rectangle{
		X:      gunPos.X,
		Y:      gunPos.Y,
		Width:  float32(gunTexture.Width) * gunScale,
		Height: float32(gunTexture.Height) * gunScale,
	}

	sourceRec := rl.NewRectangle(0, 0, float32(gunTexture.Width), float32(gunTexture.Height))
	origin := rl.Vector2{
		X: (float32(gunTexture.Width) * gunScale) / 2,
		Y: (float32(gunTexture.Height) * gunScale) / 2,
	}
	rotation := (p.Angle * 180 / math.Pi) - 90
	rl.DrawTexturePro(gunTexture, sourceRec, destRec, origin, rotation, rl.White)

	// Draw body border
	rl.DrawCircleV(p.Center, p.Radius+borderThickness, rl.Black)
	// Draw body
	rl.DrawCircleV(p.Center, p.Radius, color)

	// Draw hands
	handRadius := p.Radius * 0.35
	handDist := p.Radius * 0.85

	// Left hand offset
	leftAngle := p.Angle - math.Pi/6
	leftHandPos := rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(leftAngle)))*handDist,
		Y: p.Center.Y + float32(math.Sin(float64(leftAngle)))*handDist,
	}

	// Right hand offset
	rightAngle := p.Angle + math.Pi/6
	rightHandPos := rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(rightAngle)))*handDist,
		Y: p.Center.Y + float32(math.Sin(float64(rightAngle)))*handDist,
	}

	// Draw hand borders
	rl.DrawCircleV(leftHandPos, handRadius+borderThickness, rl.Black)
	rl.DrawCircleV(rightHandPos, handRadius+borderThickness, rl.Black)

	// Draw hands
	rl.DrawCircleV(leftHandPos, handRadius, color)
	rl.DrawCircleV(rightHandPos, handRadius, color)
}
