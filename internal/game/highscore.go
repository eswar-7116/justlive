package game

import (
	"encoding/binary"
	"log"
	"os"
	"path/filepath"
)

func highScorePath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "justlive", "highscore.bin"), nil
}

func loadHighScore() int {
	path, err := highScorePath()
	if err != nil {
		return 0
	}
	data, err := os.ReadFile(path)
	if err != nil || len(data) != 8 {
		return 0
	}
	return int(binary.LittleEndian.Uint64(data))
}

func saveHighScore(score int) {
	path, err := highScorePath()
	if err != nil {
		log.Printf("save high score: %v", err)
		return
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		log.Printf("save high score: %v", err)
		return
	}
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(score))
	if err := os.WriteFile(path, data, 0o644); err != nil {
		log.Printf("save high score: %v", err)
	}
}

func (g *Game) checkAndSaveHighScore() {
	if g.score > g.highScore {
		g.highScore = g.score
		g.newHighScore = true
		saveHighScore(g.highScore)
	}
}
