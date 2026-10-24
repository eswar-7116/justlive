package game

import (
	"image/color"

	"github.com/eswar-7116/justlive/internal/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var deathTint = color.RGBA{R: 120, G: 0, B: 0, A: 255}

func (g *Game) startDying() {
	g.checkAndSaveHighScore()
	g.gameState = GameStateDying
	g.deathTime = 0
	g.player.DeathTime = 0
	g.shake = 1
	for range deathBursts {
		g.particles = append(g.particles, particle.Burst(g.player.Center)...)
	}
}

func (g *Game) updateDying(dt float32, skip bool) {
	g.deathTime += dt
	g.player.DeathTime = g.deathTime
	slow := dt * deathTimeScale

	for _, b := range g.bullets {
		if b.Active {
			b.Update(slow)
		}
	}
	g.removeInactiveBullets()

	for _, z := range g.zombies {
		z.Update(slow, g.player)
	}
	g.resolveZombieOverlaps(slow)

	g.shake = max(g.shake-shakeDecay*dt, 0)
	g.updateEffects(slow)

	if g.deathTime >= deathDuration || skip {
		g.gameState = GameStateOver
		g.overFade = overFadeTime
	}
}

func easeInOut(t float32) float32 {
	if t < 0.5 {
		return 2 * t * t
	}
	return 1 - (-2*t+2)*(-2*t+2)/2
}

func (g *Game) drawDying() {
	e := easeInOut(min(g.deathTime/deathDuration, 1))

	// Camera slow-motion and zoom
	screenCenter := rl.Vector2{X: float32(g.width) / 2, Y: float32(g.height) / 2}
	target := rl.Vector2Lerp(screenCenter, g.player.Center, e)
	camera := rl.Camera2D{
		Offset: rl.Vector2Add(screenCenter, g.shakeOffset()),
		Target: target,
		Zoom:   1 + deathZoom*e,
	}
	g.drawWorld(camera)

	rl.DrawRectangle(0, 0, g.width, g.height, rl.Fade(rl.Black, deathDarken*e))
	rl.DrawRectangle(0, 0, g.width, g.height, rl.Fade(deathTint, 0.25*e))

	textAlpha := min(max((g.deathTime-deathTextDelay)/deathTextFade, 0), 1)
	if textAlpha > 0 {
		var fontSize int32 = 90
		text := "YOU DIED"
		x := (g.width - rl.MeasureText(text, fontSize)) / 2
		// center the camera above the body
		y := g.height / 5
		rl.DrawText(text, x, y, fontSize, rl.Fade(rl.Red, textAlpha))
	}
}
