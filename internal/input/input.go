package input

import rl "github.com/gen2brain/raylib-go/raylib"

type Input struct {
	Move  rl.Vector2
	Aim   rl.Vector2
	Shoot bool
	Dash  bool
}
