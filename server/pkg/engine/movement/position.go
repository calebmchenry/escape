package movement

type Position struct {
	// Coordinate of the southwest tile
	Coordinate Coordinate `json:"coordinate"`
	Size       int        `json:"size"`
}

func (p Position) WestEdge() []Coordinate {
	edge := make([]Coordinate, p.Size)
	x := p.Coordinate.X - 1
	for i := 0; i < p.Size; i++ {
		edge = append(edge, Coordinate{X: x, Y: p.Coordinate.Y + i})
	}
	return edge
}

func (p Position) EastEdge() []Coordinate {
	edge := make([]Coordinate, p.Size)
	x := p.Coordinate.X + p.Size
	for i := 0; i < p.Size; i++ {
		edge = append(edge, Coordinate{X: x, Y: p.Coordinate.Y + i})
	}
	return edge
}

func (p Position) SouthEdge() []Coordinate {
	edge := make([]Coordinate, p.Size)
	y := p.Coordinate.Y - 1
	for i := 0; i < p.Size; i++ {
		edge = append(edge, Coordinate{X: p.Coordinate.X + i, Y: y})
	}
	return edge
}

func (p Position) NorthEdge() []Coordinate {
	edge := make([]Coordinate, p.Size)
	y := p.Coordinate.Y - p.Size
	for i := 0; i < p.Size; i++ {
		edge = append(edge, Coordinate{X: p.Coordinate.X + i, Y: y})
	}
	return edge
}

func intersects(cs1 []Coordinate, cs2 []Coordinate) bool {
	set := make(map[string]bool)
	result := []Coordinate{}

	// Add all elements of nums1 to the set
	for _, c := range cs1 {
		set[c.String()] = true
	}

	// Check if elements of nums2 are present in the set
	for _, c := range cs2 {
		if set[c.String()] {
			result = append(result, c)
			set[c.String()] = false // To avoid duplicates in result
		}
	}

	return len(result) > 0
}

func (p Position) Orthogonal(compare Position) bool {
	return (intersects(p.NorthEdge(), compare.SouthEdge()) ||
		intersects(p.SouthEdge(), compare.NorthEdge()) ||
		intersects(p.WestEdge(), compare.EastEdge()) ||
		intersects(p.EastEdge(), compare.WestEdge()))
}

func (p Position) SouthWest() Coordinate {
	return Coordinate{X: p.Coordinate.X - 1, Y: p.Coordinate.Y - 1}
}

func (p Position) SouthEast() Coordinate {
	return Coordinate{X: p.Coordinate.X + p.Size, Y: p.Coordinate.Y - 1}
}

func (p Position) NorthWest() Coordinate {
	return Coordinate{X: p.Coordinate.X - 1, Y: p.Coordinate.Y + p.Size}
}

func (p Position) NorthEast() Coordinate {
	return Coordinate{X: p.Coordinate.X + p.Size, Y: p.Coordinate.Y + p.Size}
}

func (p Position) Diagonal(compare Position) bool {
	return (p.SouthWest().Equals(compare.NorthEast()) ||
		p.SouthEast().Equals(compare.NorthWest()) ||
		p.NorthWest().Equals(compare.SouthEast()) ||
		p.NorthEast().Equals(compare.SouthWest()))
}
