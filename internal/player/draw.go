package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) Draw(gunTexture rl.Texture2D) {
	if p.Health <= 0 {
		p.drawDead(gunTexture)
		return
	}

	color := p.Color
	if p.InvulnerableTimer > 0 {
		if int(p.InvulnerableTimer*15)%2 == 0 {
			color = rl.Red
		}
	}

	borderThickness := float32(2.5)

	// Draw dash afterimages
	for _, g := range p.ghosts {
		ghostColor := p.Color
		ghostColor.A = uint8(120 * g.life / ghostLife)
		rl.DrawCircleV(g.pos, p.Radius, ghostColor)
	}

	// Draw hands
	handRadius := p.Radius * 0.35
	handDist := p.Radius*0.85 + 10

	// Left hand offset
	leftAngle := p.Angle - math.Pi/12
	leftHandPos := rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(leftAngle)))*(handDist+15),
		Y: p.Center.Y + float32(math.Sin(float64(leftAngle)))*(handDist+15),
	}

	// Right hand offset
	rightAngle := p.Angle + math.Pi/10
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

	// Draw Gun
	gunDist := p.Radius * gunDistFactor
	gunPos := rl.Vector2{
		X: p.Center.X + float32(math.Cos(float64(p.Angle)))*gunDist,
		Y: p.Center.Y + float32(math.Sin(float64(p.Angle)))*gunDist,
	}

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
}

func easeOut(t float32) float32 {
	return 1 - (1-t)*(1-t)
}

func (p *Player) drawDead(gunTexture rl.Texture2D) {
	fall := min(p.DeathTime/deathFallTime, 1)
	settle := easeOut(fall)
	borderThickness := float32(2.5)
	dir := rl.Vector2{X: float32(math.Cos(float64(p.Angle))), Y: float32(math.Sin(float64(p.Angle)))}

	// Blood pool spread
	pool := easeOut(min(p.DeathTime/poolSpreadTime, 1))
	if pool > 0 {
		rl.DrawCircleV(p.Center, p.Radius*poolSize*pool, poolColor)
	}

	// Gun falls
	gunPos := rl.Vector2Add(p.Center, rl.Vector2Scale(dir, p.Radius*gunDistFactor+gunSlide*settle))
	destRec := rl.Rectangle{
		X:      gunPos.X,
		Y:      gunPos.Y,
		Width:  float32(gunTexture.Width) * gunScale,
		Height: float32(gunTexture.Height) * gunScale,
	}
	sourceRec := rl.NewRectangle(0, 0, float32(gunTexture.Width), float32(gunTexture.Height))
	origin := rl.Vector2{X: destRec.Width / 2, Y: destRec.Height / 2}
	gunAngle := p.Angle + gunSpin*settle
	rl.DrawTexturePro(gunTexture, sourceRec, destRec, origin, gunAngle*180/math.Pi-90, rl.White)

	color := rl.ColorBrightness(p.Color, -deathDarken*settle)

	// Hands fall
	handRadius := p.Radius * 0.35
	handDist := p.Radius*0.85 + 10
	for _, side := range []float32{-1, 1} {
		angle := float64(p.Angle + side*(math.Pi/10+limbSpread*settle))
		hand := rl.Vector2{
			X: p.Center.X + float32(math.Cos(angle))*handDist,
			Y: p.Center.Y + float32(math.Sin(angle))*handDist,
		}
		rl.DrawCircleV(hand, handRadius+borderThickness, rl.Black)
		rl.DrawCircleV(hand, handRadius, color)
	}

	// Body bulges
	radius := p.Radius * (1 + deathSquash*float32(math.Sin(float64(fall)*math.Pi)))
	rl.DrawCircleV(p.Center, radius+borderThickness, rl.Black)
	rl.DrawCircleV(p.Center, radius, color)
}
