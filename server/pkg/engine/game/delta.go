package game

import "webl-fun/pkg/engine/movement"

type Ticker interface {
	Tick(delta *Delta)
}

type Delta struct {
	Characters []*CharacterDelta `json:"characters"`
}

type CharacterDelta struct {
	CharacterID string               `json:"characterId"`
	Position    *movement.Coordinate `json:"position"`
}
