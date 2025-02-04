package game

import "webl-fun/pkg/engine/movement"

func (g *Game) Move(pID string, coord movement.Coordinate) {
	// Get player
	// Set intent to move
}

func (g *Game) Eat(pID string, targetID string) {
	// Get player's inventory
	// Get targets item type
	// Validate can eat
	// Queue eating for processing
}

func (g *Game) Attack(pID string, targetID string) {

}

func (g *Game) TalkTo(pID string, targetID string) {

}

func (g *Game) Timber(pID string, targetID string) {

}

func (g *Game) Mine(pID string, targetID string) {

}

func (g *Game) Fish(pID string, targetID string) {

}

func (g *Game) Examine(pID string, targetID string) {

}
