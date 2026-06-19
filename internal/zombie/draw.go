package zombie

import rl "github.com/gen2brain/raylib-go/raylib"

func (z *Zombie) Draw(textures []rl.Texture2D) {
	if len(textures) == 0 {
		rl.DrawRectangleV(z.Position, z.Size, z.Color)
		return
	}
	texture := textures[z.animFrame]
	sourceRec := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	if z.direction.X < 0 {
		sourceRec.Width = -sourceRec.Width
	}
	drawRec := z.DrawRec()
	origin := rl.NewVector2(0, 0)
	rl.DrawTexturePro(texture, sourceRec, drawRec, origin, 0, rl.White)
}
