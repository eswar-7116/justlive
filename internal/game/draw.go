package game

import (
	"fmt"
	"image/color"

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
}

func (g *Game) drawGameplay() {
	g.player.Draw(g.gunTexture)

	for _, b := range g.bullets {
		b.Draw(g.bulletTexture)
	}

	for _, z := range g.zombies {
		z.Draw(g.zombieTextures)
	}

	g.drawUI()
}

func (g *Game) drawUI() {
	barWidth := float32(120)
	barHeight := float32(5)
	barX := float32(42)
	barY := float32(23)

	rl.DrawText("HP", 20, 20, 12, rl.LightGray)

	bgBarRec := rl.NewRectangle(barX, barY, barWidth, barHeight)
	rl.DrawRectangleRounded(bgBarRec, 0.5, 4, color.RGBA{R: 50, G: 50, B: 50, A: 120})

	healthPct := g.player.Health / 100.0
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

	hpText := fmt.Sprintf("%.0f%%", g.player.Health)
	if g.player.Health < 0 {
		hpText = "0%"
	}
	rl.DrawText(hpText, int32(barX+barWidth+8), 20, 12, rl.LightGray)

	scoreText := fmt.Sprintf("Score: %d", g.score)
	rl.DrawText(scoreText, g.width - rl.MeasureText(scoreText, 20) - 20, 20, 20, rl.RayWhite)
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

	var pFontSize int32 = 30
	pText := "Press P to play"
	pWidth := rl.MeasureText(pText, pFontSize)
	pX := (g.width - pWidth) / 2
	pY := scoreY + scoreFontSize + 40

	rl.DrawText(pText, pX, pY, pFontSize, rl.RayWhite)
}
