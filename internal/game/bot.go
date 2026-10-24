package game

import (
	"github.com/eswar-7116/justlive/internal/bot"
	"github.com/eswar-7116/justlive/internal/drop"
)

const botAddr = "0.0.0.0:7777"

var dropKindNames = map[drop.Kind]string{
	drop.KindMedkit:    "medkit",
	drop.KindMaxHPBuff: "maxHpBuff",
	drop.KindKnockback: "knockback",
}

var gameStateNames = map[GameState]string{
	GameStateMenu:    "menu",
	GameStatePlaying: "playing",
	GameStateDying:   "dying",
	GameStateOver:    "over",
}

func (g *Game) botState() bot.State {
	p := g.player
	st := bot.State{
		State:  gameStateNames[g.gameState],
		Width:  g.width,
		Height: g.height,
		Score:  g.score,
		Player: bot.PlayerState{
			X:              p.Center.X,
			Y:              p.Center.Y,
			Radius:         p.Radius,
			HP:             p.Health,
			MaxHP:          p.MaxHealth,
			Angle:          p.Angle,
			DashReady:      p.DashCooldown <= 0 && !p.IsDashing(),
			KnockbackBoost: max(p.KnockbackBoost, 0),
		},
		Zombies: make([]bot.ZombieState, 0, len(g.zombies)),
		Drops:   make([]bot.DropState, 0, len(g.drops)),
	}

	for _, z := range g.zombies {
		c := z.Center()
		st.Zombies = append(st.Zombies, bot.ZombieState{X: c.X, Y: c.Y, HP: z.Health})
	}
	for _, d := range g.drops {
		st.Drops = append(st.Drops, bot.DropState{X: d.Position.X, Y: d.Position.Y, Kind: dropKindNames[d.Kind]})
	}
	return st
}
