package drop

import (
	"math/rand"

	"github.com/eswar-7116/justlive/internal/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Kind int

const (
	KindMedkit Kind = iota
	KindMaxHPBuff
	KindKnockback
)

const (
	Radius             float32 = 12 // px
	dropLifetime       float32 = 10 // seconds
	blinkTime          float32 = 2  // seconds before despawn
	medkitHeal         float32 = 20 // HP
	maxHPBuff          float32 = 10 // HP
	knockbackBoostTime float32 = 5  // seconds
)

type Drop struct {
	Position rl.Vector2
	Kind     Kind
	Life     float32
	Active   bool
}

func RandomDrop(pos rl.Vector2) *Drop {
	// medkit 50%, max HP buff 25%, knockback 25%
	kind := KindMedkit
	r := rand.Float32()
	if r >= 0.75 {
		kind = KindKnockback
	} else if r >= 0.5 {
		kind = KindMaxHPBuff
	}

	return &Drop{
		Position: pos,
		Kind:     kind,
		Life:     dropLifetime,
		Active:   true,
	}
}

// Returns false if the drop is not picked up
func (d *Drop) Apply(p *player.Player) bool {
	switch d.Kind {
	case KindMedkit:
		if p.Health >= p.MaxHealth {
			return false
		}
		p.Health = min(p.Health+medkitHeal, p.MaxHealth)
	case KindMaxHPBuff:
		p.MaxHealth += maxHPBuff
	case KindKnockback:
		p.KnockbackBoost = knockbackBoostTime
	}
	return true
}
