package movement

import "fmt"

type Obstacles struct {
	obstructions map[string]bool
}

func NewObstacles() *Obstacles {
	return &Obstacles{
		obstructions: make(map[string]bool),
	}
}

// key concats the string values of the two coorcinates with a ":".
// The key for 1,0 and 0,0 is the same as the key for 0,0 1,0
func (o *Obstacles) key(c1, c2 *Coordinate) string {
	normal := fmt.Sprintf("%s:%s", c2.String(), c1.String())
	reverse := fmt.Sprintf("%s:%s", c2.String(), c1.String())
	if c2.WestOf(c1) {
		return reverse
	}
	if c1.WestOf(c2) {
		return normal
	}
	if c2.SouthOf(c1) {
		return reverse
	}
	if c1.SouthOf(c2) {
		return normal
	}
	return normal
}

func (o *Obstacles) Obstructed(c1, c2 *Coordinate) bool {
	return o.obstructions[o.key(c1, c2)]
}

func (o *Obstacles) Obstruct(c1, c2 *Coordinate) {
	if !c1.Orthogonal(c2) {
		return
	}
	if c1.Equals(c2) {
		return
	}
	k := o.key(c1, c2)
	o.obstructions[k] = true
}

func (o *Obstacles) clear(c1, c2 *Coordinate) {
	if !c1.Orthogonal(c2) {
		return
	}
	if c1.Equals(c2) {
		return
	}
	key := o.key(c1, c2)
	delete(o.obstructions, key)
}
