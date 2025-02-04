package game

import "webl-fun/pkg/engine/movement"

type Player struct {
	ID       string
	position movement.Coordinate
	path     []movement.Coordinate
}

func NewPlayer(id string, c movement.Coordinate) *Player {
	return &Player{ID: id, position: c}
}

func (p *Player) MoveTo(dst movement.Coordinate) {
	o := movement.NewObstacles()
	p.path = movement.Path(o, p.position, dst)
}

func (p *Player) Tick() *EntityDelta {
	if len(p.path) > 0 {
		p.position = p.path[0]
		p.path = p.path[1:]
	}
	d := &EntityDelta{
		ID:   p.ID,
		Type: "Player",
		Position: Position{
			X:    p.position.X,
			Y:    p.position.Y,
			Size: 1,
		},
		Animation: "move",
		Actions:   []Action{},
	}
	return d
}
