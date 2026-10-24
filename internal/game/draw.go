package game

import (
	"fmt"
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) draw() {
	switch g.gameState {
	case GameStateMenu:
		g.drawMenu()
	case GameStatePlaying:
		g.drawGameplay()
	case GameStateOver:
		g.drawGameover()
	}
}

func (g *Game) drawMenu() {
	var titleFontSize int32 = 100
	var textWidth int32 = rl.MeasureText(g.title, titleFontSize)
	var titleX int32 = (g.width - textWidth) / 2
	var titleY int32 = (g.height - titleFontSize) / 2
	rl.DrawText(g.title, titleX, titleY, titleFontSize, rl.Red)

	var fontSize int32 = 30
	text := "Press P to start surviving"
	textWidth = rl.MeasureText(text, fontSize)
	var X int32 = (g.width - textWidth) / 2
	var Y int32 = ((g.height - fontSize) / 2) + titleFontSize
	rl.DrawText(text, X, Y, fontSize, rl.RayWhite)

	if g.highScore > 0 {
		var hsFontSize int32 = 20
		hsText := fmt.Sprintf("High Score: %d", g.highScore)
		hsX := (g.width - rl.MeasureText(hsText, hsFontSize)) / 2
		rl.DrawText(hsText, hsX, Y+fontSize+20, hsFontSize, rl.LightGray)
	}
}

func (g *Game) drawGameplay() {
	rl.BeginMode2D(rl.Camera2D{Offset: g.shakeOffset(), Zoom: 1})

	// Shells on the ground lie under everything
	for _, s := range g.shells {
		if s.Grounded {
			s.Draw()
		}
	}

	for _, d := range g.drops {
		d.Draw()
	}

	g.player.Draw(g.gunTexture)

	// Shells in the air fly over the player
	for _, s := range g.shells {
		if !s.Grounded {
			s.Draw()
		}
	}

	for _, b := range g.bullets {
		b.Draw(g.bulletTexture)
	}

	for _, z := range g.zombies {
		z.Draw(g.zombieTextures)
	}

	for _, p := range g.particles {
		p.Draw()
	}

	rl.EndMode2D()

	g.drawUI()
}

func (g *Game) shakeOffset() rl.Vector2 {
	strength := maxShakeOffset * g.shake * g.shake
	return rl.Vector2{
		X: (rand.Float32()*2 - 1) * strength,
		Y: (rand.Float32()*2 - 1) * strength,
	}
}

func (g *Game) drawUI() {
	barWidth := float32(120)
	barHeight := float32(5)
	barX := float32(42)
	barY := float32(23)

	rl.DrawText("HP", 20, 20, 12, rl.LightGray)

	bgBarRec := rl.NewRectangle(barX, barY, barWidth, barHeight)
	rl.DrawRectangleRounded(bgBarRec, 0.5, 4, color.RGBA{R: 50, G: 50, B: 50, A: 120})

	healthPct := g.player.Health / g.player.MaxHealth
	if healthPct < 0 {
		healthPct = 0
	}
	if healthPct > 1 {
		healthPct = 1
	}

	var barColor color.RGBA
	if healthPct > 0.5 {
		barColor = color.RGBA{R: 46, G: 204, B: 113, A: 200} // Soft Green
	} else if healthPct > 0.2 {
		barColor = color.RGBA{R: 241, G: 196, B: 15, A: 200} // Soft Yellow/Orange
	} else {
		barColor = color.RGBA{R: 231, G: 76, B: 60, A: 200} // Soft Red
	}

	if healthPct > 0 {
		fillRec := rl.NewRectangle(barX, barY, barWidth*healthPct, barHeight)
		rl.DrawRectangleRounded(fillRec, 0.5, 4, barColor)
	}

	rl.DrawRectangleRoundedLines(bgBarRec, 0.5, 4, color.RGBA{R: 120, G: 120, B: 120, A: 100})

	hpText := fmt.Sprintf("%.0f/%.0f", max(g.player.Health, 0), g.player.MaxHealth)
	rl.DrawText(hpText, int32(barX+barWidth+8), 20, 12, rl.LightGray)

	if g.player.KnockbackBoost > 0 {
		boostText := fmt.Sprintf("Knockback %.1fs", g.player.KnockbackBoost)
		rl.DrawText(boostText, 20, 38, 12, color.RGBA{R: 230, G: 126, B: 34, A: 255})
	}

	scoreText := fmt.Sprintf("Score: %d", g.score)
	rl.DrawText(scoreText, g.width - rl.MeasureText(scoreText, 20) - 20, 20, 20, rl.RayWhite)

	bestText := fmt.Sprintf("Best: %d", max(g.highScore, g.score))
	rl.DrawText(bestText, g.width-rl.MeasureText(bestText, 12)-20, 44, 12, rl.LightGray)
}

func (g *Game) drawGameover() {
	var titleFontSize int32 = 80
	titleText := "Game Over"
	titleWidth := rl.MeasureText(titleText, titleFontSize)
	titleX := (g.width - titleWidth) / 2
	titleY := (g.height - titleFontSize) / 2 - 60

	rl.DrawText(titleText, titleX, titleY, titleFontSize, rl.Red)

	var scoreFontSize int32 = 40
	scoreText := fmt.Sprintf("Final Score: %d", g.score)
	scoreWidth := rl.MeasureText(scoreText, scoreFontSize)
	scoreX := (g.width - scoreWidth) / 2
	scoreY := titleY + titleFontSize + 20

	rl.DrawText(scoreText, scoreX, scoreY, scoreFontSize, rl.LightGray)

	var hsFontSize int32 = 24
	hsText := fmt.Sprintf("High Score: %d", g.highScore)
	hsColor := rl.LightGray
	if g.newHighScore {
		hsText = "New High Score!"
		hsColor = rl.Yellow
	}
	hsX := (g.width - rl.MeasureText(hsText, hsFontSize)) / 2
	hsY := scoreY + scoreFontSize + 16
	rl.DrawText(hsText, hsX, hsY, hsFontSize, hsColor)

	var pFontSize int32 = 30
	pText := "Press P to play"
	pWidth := rl.MeasureText(pText, pFontSize)
	pX := (g.width - pWidth) / 2
	pY := hsY + hsFontSize + 30

	rl.DrawText(pText, pX, pY, pFontSize, rl.RayWhite)
}
