package movement

import (
	"fmt"
	"math"
)

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coordinate) West() Coordinate {
	return Coordinate{X: c.X - 1, Y: c.Y}
}

func (c Coordinate) East() Coordinate {
	return Coordinate{X: c.X + 1, Y: c.Y}
}

func (c Coordinate) South() Coordinate {
	return Coordinate{X: c.X, Y: c.Y - 1}
}

func (c Coordinate) North() Coordinate {
	return Coordinate{X: c.X, Y: c.Y + 1}
}

func (c Coordinate) Orthogonal(coord Coordinate) bool {
	dx := c.X - coord.X
	dy := c.Y - coord.Y
	return (math.Abs(float64(dx)) + math.Abs(float64(dy))) == 1
}

func (c Coordinate) WestOf(coord Coordinate) bool {
	return c.X < coord.X
}

func (c Coordinate) SouthOf(coord Coordinate) bool {
	return c.Y < coord.Y
}

func (c Coordinate) SouthWest() Coordinate {
	return Coordinate{X: c.X - 1, Y: c.Y - 1}
}

func (c Coordinate) SouthEast() Coordinate {
	return Coordinate{X: c.X + 1, Y: c.Y - 1}
}

func (c Coordinate) NorthWest() Coordinate {
	return Coordinate{X: c.X - 1, Y: c.Y + 1}
}

func (c Coordinate) NorthEast() Coordinate {
	return Coordinate{X: c.X + 1, Y: c.Y + 1}
}

func (c Coordinate) Diagonal(coord Coordinate) bool {
	dx := c.X - coord.X
	dy := c.Y - coord.Y
	return (math.Abs(float64(dx)) == 1) && (math.Abs(float64(dy)) == 1)
}

func (c Coordinate) Equals(coord Coordinate) bool {
	return c.X == coord.X && c.Y == coord.Y
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
