package game

import "webl-fun/pkg/engine/movement"

type Character struct {
	ID       string
	position *movement.Coordinate
	path     []*movement.Coordinate
}

var _ Ticker = &Character{}

func NewCharacter(id string, c *movement.Coordinate) *Character {
	return &Character{ID: id, position: c}
}

func (c *Character) MoveTo(dst *movement.Coordinate) {
	o := movement.NewObstacles()
	c.path = movement.Path(o, c.position, dst)
}

func (c *Character) Tick(delta *Delta) {
	if len(c.path) > 0 {
		c.position = c.path[0]
		c.path = c.path[1:]
	}
	dC := &CharacterDelta{CharacterID: c.ID, Position: c.position}
	delta.Characters = append(delta.Characters, dC)
}
