package game

import (
	"fmt"
	"log"

	"github.com/eswar-7116/justlive/assets"
	"github.com/eswar-7116/justlive/internal/bullet"
	"github.com/eswar-7116/justlive/internal/drop"
	"github.com/eswar-7116/justlive/internal/particle"
	"github.com/eswar-7116/justlive/internal/player"
	"github.com/eswar-7116/justlive/internal/zombie"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const maxFrameTime float32 = 0.05 // 20 FPS floor

// Player
const (
	playerRadius float32 = 30
	baseFireRate float32 = 0.15
)

// Zombie
const (
	zombieSize      float32 = 30
	zombieHealth    float32 = 100
	zombieJostle    float32 = 120 // extra push px/s, 2px at 60 FPS
	bulletKnockback float32 = 100 // px/s
)

// Spawning
const (
	startSpawnInterval float32 = 2.0
	minSpawnInterval   float32 = 0.4
	spawnIntervalDecay float32 = 0.97
)

// Difficulty
const (
	baseZombieSpeed float32 = 50
	maxZombieSpeed  float32 = 200 // stay below player speed
	zombieSpeedStep float32 = 10
	killScore               = 10
	speedUpScore            = 50 // speed up zombies every 50 score
)

// Screen shake
const (
	shakePerHit    float32 = 0.6
	shakeDecay     float32 = 2.5 // per second
	maxShakeOffset float32 = 25  // px
)

// Drops
const (
	dropChance         float32 = 0.15 // per kill
	knockbackBoostMult float32 = 4
)

type Game struct {
	// Window
	width     int32
	height    int32
	title     string
	gameState GameState

	// Textures
	zombieTextures []rl.Texture2D
	gunTexture     rl.Texture2D
	bulletTexture  rl.Texture2D

	// Player
	player     *player.Player
	shootTimer float32
	fireRate   float32

	// Entities
	zombies []*zombie.Zombie
	bullets []*bullet.Bullet
	drops   []*drop.Drop

	// Spawning and difficulty
	spawnTimer    float32
	spawnInterval float32
	zombieSpeed   float32

	// Score
	score        int
	highScore    int
	newHighScore bool

	// Effects
	shake     float32
	particles []*particle.Particle
}

func NewGame(width, height int32, title string) *Game {
	g := &Game{
		width:     width,
		height:    height,
		title:     title,
		gameState: GameStateMenu,
		fireRate:  baseFireRate,
		highScore: loadHighScore(),
	}
	g.reset()
	return g
}

func (g *Game) reset() {
	// Player
	center := rl.Vector2{X: float32(g.width) / 2, Y: float32(g.height) / 2}
	g.player = player.NewPlayer(center, playerRadius)
	g.shootTimer = 0

	// Entities
	g.zombies = g.zombies[:0]
	g.bullets = g.bullets[:0]
	g.drops = g.drops[:0]

	// Spawning and difficulty
	g.spawnTimer = 0
	g.spawnInterval = startSpawnInterval
	g.zombieSpeed = baseZombieSpeed

	// Score
	g.score = 0
	g.newHighScore = false

	// Effects
	g.shake = 0
	g.particles = g.particles[:0]
}

func (g *Game) Run() {
	rl.InitWindow(int32(g.width), int32(g.height), g.title)
	defer rl.CloseWindow()

	// Disable the exit key (ESC) to prevent the game from closing
	rl.SetExitKey(0)

	// Load zombie textures
	g.zombieTextures = make([]rl.Texture2D, 10)
	for i := range 10 {
		g.zombieTextures[i] = loadTexture(fmt.Sprintf("zombie%d.png", i+1))
	}
	g.gunTexture = loadTexture("gun.png")
	g.bulletTexture = loadTexture("bullet.png")
	defer func() {
		for _, tex := range g.zombieTextures {
			rl.UnloadTexture(tex)
		}
		rl.UnloadTexture(g.gunTexture)
		rl.UnloadTexture(g.bulletTexture)
	}()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		dt := min(rl.GetFrameTime(), maxFrameTime)
		g.update(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		g.draw()
		rl.EndDrawing()
	}

	// Window closed mid-run
	if g.gameState == GameStatePlaying {
		g.checkAndSaveHighScore()
	}
}

func loadTexture(name string) rl.Texture2D {
	data, err := assets.FS.ReadFile(name)
	if err != nil {
		log.Fatalf("load asset %s: %v", name, err)
	}
	img := rl.LoadImageFromMemory(".png", data, int32(len(data)))
	defer rl.UnloadImage(img)
	return rl.LoadTextureFromImage(img)
}
