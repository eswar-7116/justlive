package game

import rl "github.com/gen2brain/raylib-go/raylib"

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
	g.player.Draw()

	for _, z := range g.zombies {
		z.Draw()
	}
}

func (g *Game) drawGameover() {
	rl.DrawText("Game Over", 100, 100, 50, rl.White)
}
