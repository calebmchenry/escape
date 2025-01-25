package movement

type state struct {
	coord *Coordinate
	path  []*Coordinate
}

func makeState(p []*Coordinate, t *Coordinate) state {
	newPath := make([]*Coordinate, len(p))
	copy(newPath, p)
	newPath = append(newPath, t)
	return state{
		coord: t,
		path:  newPath,
	}
}

func Path(o *Obstacles, start *Coordinate, destination *Coordinate) []*Coordinate {
	queue := []state{{coord: start, path: []*Coordinate{}}}
	processed := make(map[string]bool)

	for len(queue) != 0 {
		value := queue[0]
		queue = queue[1:]
		tile := value.coord

		if processed[tile.String()] {
			continue
		}
		processed[tile.String()] = true

		if tile.Equals(destination) {
			return value.path
		}

		orthogonals := []*Coordinate{
			tile.West(),
			tile.East(),
			tile.South(),
			tile.North(),
		}

		for _, t := range orthogonals {
			if !o.Obstructed(tile, t) {
				queue = append(queue, makeState(value.path, t))
			}
		}

		w := o.Obstructed(tile, tile.West())
		e := o.Obstructed(tile, tile.East())
		s := o.Obstructed(tile, tile.South())
		n := o.Obstructed(tile, tile.North())

		ws := o.Obstructed(tile.West(), tile.SouthWest())
		wn := o.Obstructed(tile.West(), tile.NorthWest())
		es := o.Obstructed(tile.East(), tile.SouthEast())
		en := o.Obstructed(tile.East(), tile.NorthEast())
		sw := o.Obstructed(tile.South(), tile.SouthWest())
		se := o.Obstructed(tile.South(), tile.SouthEast())
		nw := o.Obstructed(tile.North(), tile.NorthWest())
		ne := o.Obstructed(tile.North(), tile.NorthEast())

		if (!w && !ws) || (!s && !sw) {
			queue = append(queue, makeState(value.path, tile.SouthWest()))
		}
		if (!e && !es) || (!s && !se) {
			queue = append(queue, makeState(value.path, tile.SouthEast()))
		}
		if (!w && !wn) || (!n && !nw) {
			queue = append(queue, makeState(value.path, tile.NorthWest()))
		}
		if (!e && !en) || (!n && !ne) {
			queue = append(queue, makeState(value.path, tile.NorthEast()))
		}
	}
	return []*Coordinate{}
}
