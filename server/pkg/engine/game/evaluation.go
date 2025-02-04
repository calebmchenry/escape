package game

func (g *Game) evaluateTick() {
	// Instant inventory actions (equip, drop, etc)
	//   Perhaps this can be done at command time
	// Eat
	// Apply damage buffers
	// Move npcs
	//   If target is moving to tile then move to future position
	//   If target is moving to a target then move to current position
	//     Or maybe moving to future tile as if target were a static tile (even if moving)
	// Move players
	//   If moving target determine targets next location
	//     How would I handle a chain
	// Determine queued action
	// Execute queued action
}
