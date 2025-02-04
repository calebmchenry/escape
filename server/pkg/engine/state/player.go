package state

import (
	"fmt"
	"webl-fun/pkg/engine/movement"
)

type Levels struct {
	// combat
	Melee     int
	Ranged    int
	Magic     int
	Hitpoints int
	// Radiance  int

	// support
	Agility int

	// gathering
	// Mining      int
	// Woodcutting int
	// Artificing  int
	Fishing int
	// Farming     int
	// Foraging    int
	// Hunting     int

	// processing
	Cooking int
	// Alchemy   int
	// Forging   int
	// Fletching int
	// Arcana    int
	// Crafting  int
}

type Equipment struct {
	ID   string
	Name string

	MeleeAccuracy  int
	RangedAccuracy int
	MagicAccuracy  int

	MeleeDamage  int
	RangedDamage int
	MagicDamage  int

	MeleeDefense  int
	RangedDefense int
	MagicDefense  int
}

type Equipped struct {
	Head   Equipment
	Body   Equipment
	Gloves Equipment
	Legs   Equipment
	Boots  Equipment

	Back Equipment
	Ring Equipment
	Neck Equipment
}

type Splat struct {
	SrcID string
	Value int
}

type TickBuffer[T any] struct {
	currentTick int
	buffer      map[int][]T
}

func (b *TickBuffer[T]) Add(value T, ticks int) error {
	if ticks < 1 {
		return fmt.Errorf("cannot add hits %d tick(s) away", ticks)
	}
	key := b.currentTick + ticks
	b.buffer[key] = append(b.buffer[key], value)
	return nil
}

func (b *TickBuffer[T]) Tick() []T {
	values := b.buffer[b.currentTick]
	delete(b.buffer, b.currentTick)
	b.currentTick += 1
	return values
}

func (b *TickBuffer[T]) Clear() {
	for k := range b.buffer {
		delete(b.buffer, k)
	}
}

type Player struct {
	TrueTile     *movement.Coordinate
	Path         []*movement.Coordinate
	DamageBuffer TickBuffer[Splat]
	HealBuffer   TickBuffer[Splat]
	TargetID     string
	Running      bool
	Gear         *Equipped
	Inventory    *Inventory
}

func (p *Player) Tick() {
	// Progress path
	// Progress damage buffers
	// Progress Heal buffers
}
