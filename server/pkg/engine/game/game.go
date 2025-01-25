package game

import (
	"encoding/json"
	"fmt"
	"time"
	"webl-fun/pkg/engine/movement"
)

type Game struct {
	characters map[string]*Character
	subcribers map[string]chan *Tick
}

func New() *Game {
	cs := make(map[string]*Character)
	g := &Game{characters: cs, subcribers: make(map[string]chan *Tick)}
	return g
}

var TICK = 600 * time.Millisecond

type Tick struct {
	Tick  int    `json:"tick"`
	Delta *Delta `json:"delta"`
}

func (g *Game) Start() {
	ticker := time.NewTicker(TICK)
	quit := make(chan struct{})

	go func() {
		count := 0
		for {
			select {
			case <-ticker.C:
				// update everything
				d := g.update()

				t := &Tick{
					Tick:  count,
					Delta: d,
				}
				count++
				// send to connections
				fmt.Printf("Sending tick %d to %d connection(s)\n", count, len(g.subcribers))
				for _, ch := range g.subcribers {
					ch <- t
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (g *Game) AddCharacter(c *Character) {
	g.characters[c.ID] = c
}

func (g *Game) Subscribe(cID string) <-chan *Tick {
	dst := make(chan *Tick)
	g.AddCharacter(NewCharacter(cID, &movement.Coordinate{X: 0, Y: 0}))
	g.subcribers[cID] = dst
	return dst
}

func (g *Game) Unsubscribe(cID string) {
	dst := g.subcribers[cID]
	close(dst)
	delete(g.subcribers, cID)
	delete(g.characters, cID)
}

func (g *Game) Act(characterID string, actionID string) {
	c := g.characters[characterID]
	b := []byte(actionID)
	var coord movement.Coordinate
	err := json.Unmarshal(b, &coord)
	if err != nil {
		fmt.Printf("Failed to parse coordinate: %s\n", actionID)
	}
	c.MoveTo(&coord)
}

func (g *Game) update() *Delta {
	d := &Delta{}
	for _, c := range g.characters {
		c.Tick(d)
	}
	return d
}
