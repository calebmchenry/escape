package game

// Inventory
func (g *Game) dropItem(pID string, targetID string) {
	// Get player's inventory
	// Remove item from inventory
	// Place one ground
	// Queue to deltas
}

func (g *Game) swapItems(pID string, srcID string, targetID string) {
	// Get player's inventory
	// Swap items
	// Queue item deltas
}

func (g *Game) equipItem(pID string, targetID string) {
	// Get player's inventory
	// Confirm item can be equipped
	// Add to equipment
	// Replace same position with previously equipped item
}

func (g *Game) useItems(pID string, srcID string, targetID string) {
	// Get player's inventory
	// Get source item type
	// Get target item type
	// Resolve usage
	// Queue item deltas
}
