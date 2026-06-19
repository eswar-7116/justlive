package game

import (
	"github.com/eswar-7116/justlive/internal/player"
	"github.com/eswar-7116/justlive/internal/zombie"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width     int32
	height    int32
	title     string
	gameState GameState
	player    *player.Player
	zombies   []*zombie.Zombie

	spawnTimer    float32
	spawnInterval float32
}

func NewGame(width, height int32, title string) *Game {
	center := rl.Vector2{X: float32(width) / 2, Y: float32(height) / 2}
	p := player.NewPlayer(center, 30)

	return &Game{
		width:         width,
		height:        height,
		title:         title,
		gameState:     GameStateMenu,
		player:        p,
		zombies:       make([]*zombie.Zombie, 0),
		spawnInterval: 2,
	}
}

func (g *Game) Run() {
	rl.InitWindow(int32(g.width), int32(g.height), g.title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		g.update(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		g.draw()
		rl.EndDrawing()
	}
}
