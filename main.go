package main

import "github.com/eswar-7116/justlive/internal/game"

func main() {
	g := game.NewGame(1200, 700, "Just Live!")
	g.Run()
}
