package game

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"webl-fun/pkg/engine/movement"
)

type Game struct {
	players     map[string]*Player
	subscribers map[string]chan *ServerTick
}

func New() *Game {
	cs := make(map[string]*Player)
	g := &Game{players: cs, subscribers: make(map[string]chan *ServerTick)}
	return g
}

var TICK = 600 * time.Millisecond

func (g *Game) Start() {
	ticker := time.NewTicker(TICK)
	quit := make(chan struct{})

	go func() {
		count := 0
		for {
			select {
			case <-ticker.C:
				startedAt := time.Now().Unix()
				// update everything
				ds := g.update()

				endedAt := time.Now().Unix()
				t := &ServerTick{
					Tick:          count,
					TickStartedAt: int(startedAt),
					TickEndedAt:   int(endedAt),
					Instance:      "main",
					Toggles:       TogglesPayload{Run: false},
					Inventory:     []InventoryDelta{},
					Experience:    []ExperienceDelta{},
					Equipment:     []EquipmentDelta{},
				}
				count++
				// send to connections
				fmt.Printf("Sending tick %d to %d connection(s)\n", count, len(g.subscribers))
				for playerID, ch := range g.subscribers {
					t.SentAt = int(time.Now().Unix())
					for _, d := range ds {
						if d.ID == playerID {
							t.Player = d
						}
						t.NPCs = append(t.NPCs, d)
					}
					ch <- t
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (g *Game) AddPlayer(p *Player) {
	g.players[p.ID] = p
}

func (g *Game) Subscribe(playerID string) <-chan *ServerTick {
	dst := make(chan *ServerTick)
	g.AddPlayer(NewPlayer(playerID, movement.Coordinate{X: 0, Y: 0}))
	g.subscribers[playerID] = dst
	return dst
}

func (g *Game) Unsubscribe(playerID string) {
	dst := g.subscribers[playerID]
	close(dst)
	delete(g.subscribers, playerID)
	delete(g.players, playerID)
}

func (g *Game) Act(playerID, action, target, src string) {
	c := g.players[playerID]
	ss := strings.Split(target, ",")
	x, err := strconv.Atoi(ss[0])
	if err != nil {
		fmt.Printf("Failed to parse x coordinate: %s\n", target)
	}
	y, err := strconv.Atoi(ss[1])
	if err != nil {
		fmt.Printf("Failed to parse y coordinate: %s\n", target)
	}
	coord := movement.Coordinate{X: x, Y: y}
	c.MoveTo(coord)
}

func (g *Game) update() []*EntityDelta {
	deltas := []*EntityDelta{}
	for _, c := range g.players {
		deltas = append(deltas, c.Tick())
	}
	return deltas
}
