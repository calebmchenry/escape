package proximity

import (
	"math"
	"webl-fun/pkg/engine/movement"
)

type Positioner interface {
	position() movement.Position
}

type ProximityQuery struct {
	entitiesByChunk map[string]map[string]Positioner
}

func withinRadius(center movement.Coordinate, radius uint, coord movement.Coordinate) bool {
	dx := center.X - coord.X
	dy := center.Y - coord.Y
	return math.Abs(float64(dx)) <= float64(radius) || math.Abs(float64(dy)) <= float64(radius)
}
