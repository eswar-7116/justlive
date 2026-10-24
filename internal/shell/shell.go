package shell

import (
	"image/color"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Ejection
const (
	ejectMinSpeed = 140                // px/s
	ejectMaxSpeed = 170                // px/s
	ejectSpread   = 10 * math.Pi / 180 // radians each side
	ejectBack     = 0.4                // how much it flies backward
	startHeight   = 12                 // px above the ground
	popSpeed      = 120                // px/s upward
	minSpin       = 600                // degrees/s
	maxSpin       = 1000               // degrees/s
)

// Falling
const (
	gravity    = 600   // px/s²
	bounce     = 0.35  // vertical speed kept on bounce
	minBounce  = 20    // px/s, settles below this
	groundDrag = 0.002 // ground speed kept per second
)

// Looks
const (
	lifetime    = 2   // seconds
	zScale      = 80  // px of height that doubles the size
	fadeTime    = 0.5 // seconds
	width       = 4   // px
	height      = 2   // px
	rimWidth    = 1   // px, dark primer end
	groundAlpha = 0.6 // dimmed once landed
)

// Muted so it doesn't read as a bullet
var (
	brass = color.RGBA{R: 125, G: 105, B: 70, A: 255}
	rim   = color.RGBA{R: 60, G: 50, B: 35, A: 255}
)

type Shell struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Z        float32 // fake height above the ground
	VZ       float32
	Rotation float32
	Spin     float32
	Life     float32
	Grounded bool
}

func randRange(lo, hi float32) float32 {
	return lo + rand.Float32()*(hi-lo)
}

// Ejects to the gun's right side, a bit backward
func Eject(pos, aim rl.Vector2) *Shell {
	right := rl.Vector2{X: -aim.Y, Y: aim.X}
	dir := rl.Vector2Normalize(rl.Vector2Subtract(right, rl.Vector2Scale(aim, ejectBack)))
	dir = rl.Vector2Rotate(dir, randRange(-ejectSpread, ejectSpread))

	spin := randRange(minSpin, maxSpin)
	if rand.Intn(2) == 0 {
		spin = -spin
	}

	return &Shell{
		Position: pos,
		Velocity: rl.Vector2Scale(dir, randRange(ejectMinSpeed, ejectMaxSpeed)),
		Z:        startHeight,
		VZ:       popSpeed,
		Rotation: rand.Float32() * 360,
		Spin:     spin,
		Life:     lifetime,
	}
}
