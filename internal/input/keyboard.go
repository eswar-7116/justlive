package input

import rl "github.com/gen2brain/raylib-go/raylib"

func boolToFloat32(b bool) float32 {
	if b {
		return 1
	} else {
		return 0
	}
}

func FromKeyboard() Input {
	return Input{
		Move: rl.Vector2{
			X: boolToFloat32(rl.IsKeyDown(rl.KeyD)) - boolToFloat32(rl.IsKeyDown(rl.KeyA)),
			Y: boolToFloat32(rl.IsKeyDown(rl.KeyS)) - boolToFloat32(rl.IsKeyDown(rl.KeyW)),
		},
		Aim:   rl.GetMousePosition(),
		Shoot: rl.IsMouseButtonDown(rl.MouseLeftButton),
		Dash:  rl.IsKeyPressed(rl.KeySpace),
		Start: rl.IsKeyPressed(rl.KeyP),
	}
}
