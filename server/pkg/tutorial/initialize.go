package tutorial

import (
	"webl-fun/pkg/engine/movement"
	"webl-fun/pkg/engine/state"
	"webl-fun/pkg/ids"
)

const chunkSize = 127

type Chunk struct {
	Coordinate movement.Coordinate
	Obstacles  movement.Obstacles
	entities   []string
}

func (c *Chunk) translateCoordinate(coord movement.Coordinate) movement.Coordinate {
	return movement.Coordinate{
		X: c.Coordinate.X*chunkSize + coord.X,
		Y: c.Coordinate.Y*chunkSize + coord.Y,
	}
}

type Instance struct {
	Chunks  map[ids.Chunk]Chunk
	players map[string]state.Player
}

func Initialize() {

}
