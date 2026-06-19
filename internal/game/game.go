package game

import (
	"fmt"

	"github.com/eswar-7116/justlive/internal/bullet"
	"github.com/eswar-7116/justlive/internal/player"
	"github.com/eswar-7116/justlive/internal/zombie"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width          int32
	height         int32
	title          string
	gameState      GameState
	player         *player.Player
	zombies        []*zombie.Zombie
	spawnTimer     float32
	spawnInterval  float32
	zombieTextures []rl.Texture2D
	gunTexture     rl.Texture2D
	bullets        []*bullet.Bullet
	bulletTexture  rl.Texture2D
	shootTimer     float32
	fireRate       float32
	score          int
	zombieSpeed    float32
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
		bullets:       make([]*bullet.Bullet, 0),
		fireRate:      0.15,
		score:         0,
		zombieSpeed:   50,
	}
}

func (g *Game) Run() {
	rl.InitWindow(int32(g.width), int32(g.height), g.title)
	defer rl.CloseWindow()

	// Load zombie textures
	g.zombieTextures = make([]rl.Texture2D, 10)
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/zombie%d.png", i+1)
		g.zombieTextures[i] = rl.LoadTexture(path)
	}
	g.gunTexture = rl.LoadTexture("assets/gun.png")
	g.bulletTexture = rl.LoadTexture("assets/bullet.png")
	defer func() {
		for _, tex := range g.zombieTextures {
			rl.UnloadTexture(tex)
		}
		rl.UnloadTexture(g.gunTexture)
		rl.UnloadTexture(g.bulletTexture)
	}()

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
