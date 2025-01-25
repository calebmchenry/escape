package movement_test

import (
	"slices"
	"testing"
	"webl-fun/pkg/engine/movement"
)

func assertPathEquality(t *testing.T, want, got []*movement.Coordinate) {
	equal := slices.CompareFunc(want, got, func(a *movement.Coordinate, b *movement.Coordinate) int {
		if a.Equals(b) {
			return 0
		}
		return -1
	}) == 0
	if !equal {
		t.Errorf("Expected want: '%s' to equal got: '%s'", want, got)
	}
}

func want(coords ...*movement.Coordinate) []*movement.Coordinate {
	return coords
}

func TestPath(t *testing.T) {
	o := movement.NewObstacles()
	start := &movement.Coordinate{X: 0, Y: 0}
	got := movement.Path(o, start, start)

	if len(got) != 0 {
		t.Errorf("Expect length 0, received: '%s'", got)
	}
}

func TestPathAdjacent(t *testing.T) {
	o := movement.NewObstacles()
	start := &movement.Coordinate{X: 0, Y: 0}

	// orthogonal
	got := movement.Path(o, start, start.West())
	assertPathEquality(t, want(start.West()), got)
	got = movement.Path(o, start, start.East())
	assertPathEquality(t, want(start.East()), got)
	got = movement.Path(o, start, start.South())
	assertPathEquality(t, want(start.South()), got)
	got = movement.Path(o, start, start.North())
	assertPathEquality(t, want(start.North()), got)

	// diagonal
	got = movement.Path(o, start, start.SouthWest())
	assertPathEquality(t, want(start.SouthWest()), got)
	got = movement.Path(o, start, start.SouthEast())
	assertPathEquality(t, want(start.SouthEast()), got)
	got = movement.Path(o, start, start.NorthWest())
	assertPathEquality(t, want(start.NorthWest()), got)
	got = movement.Path(o, start, start.NorthEast())
	assertPathEquality(t, want(start.NorthEast()), got)
}

func TestPathAdjacentDistance2(t *testing.T) {
	o := movement.NewObstacles()
	start := &movement.Coordinate{X: 0, Y: 0}

	w := start.West()
	e := start.East()
	s := start.South()
	n := start.North()
	sw := start.SouthWest()
	se := start.SouthEast()
	nw := start.NorthWest()
	ne := start.NorthEast()

	// orthogonal
	ww := start.West().West()
	got := movement.Path(o, start, ww)
	assertPathEquality(t, want(w, ww), got)

	ee := start.East().East()
	got = movement.Path(o, start, ee)
	assertPathEquality(t, want(e, ee), got)

	ss := start.South().South()
	got = movement.Path(o, start, ss)
	assertPathEquality(t, want(s, ss), got)

	nn := start.North().North()
	got = movement.Path(o, start, nn)
	assertPathEquality(t, want(n, nn), got)

	// diagonal
	swsw := start.SouthWest().SouthWest()
	got = movement.Path(o, start, swsw)
	assertPathEquality(t, want(sw, swsw), got)

	sese := start.SouthEast().SouthEast()
	got = movement.Path(o, start, sese)
	assertPathEquality(t, want(se, sese), got)

	nwnw := start.NorthWest().NorthWest()
	got = movement.Path(o, start, nwnw)
	assertPathEquality(t, want(nw, nwnw), got)

	nene := start.NorthEast().NorthEast()
	got = movement.Path(o, start, nene)
	assertPathEquality(t, want(ne, nene), got)

	// L's
	wsw := start.West().SouthWest()
	got = movement.Path(o, start, wsw)
	assertPathEquality(t, want(w, wsw), got)

	ese := start.East().SouthEast()
	got = movement.Path(o, start, ese)
	assertPathEquality(t, want(e, ese), got)

	wnw := start.West().NorthWest()
	got = movement.Path(o, start, wnw)
	assertPathEquality(t, want(w, wnw), got)

	ene := start.East().NorthEast()
	got = movement.Path(o, start, ene)
	assertPathEquality(t, want(e, ene), got)

	ssw := start.South().SouthWest()
	got = movement.Path(o, start, ssw)
	assertPathEquality(t, want(s, ssw), got)

	sse := start.South().SouthEast()
	got = movement.Path(o, start, sse)
	assertPathEquality(t, want(s, sse), got)

	nnw := start.North().NorthWest()
	got = movement.Path(o, start, nnw)
	assertPathEquality(t, want(n, nnw), got)

	nne := start.North().NorthEast()
	got = movement.Path(o, start, nne)
	assertPathEquality(t, want(n, nne), got)
}
