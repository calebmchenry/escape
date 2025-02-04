package state

import (
	"errors"
)

type Item struct {
	ID       string
	Quantity int
}

type Inventory struct {
	items map[int]*Item
}

var blankItem = &Item{
	ID: blankID,
}

const blankID = "__blank"

var ErrBadInventoryIndex = errors.New("bad inventory index")
var ErrBadInventoryItemQuantity = errors.New("bad item quantity")
var ErrBadInventoryItemID = errors.New("bad item ID")
var ErrInventoryFull = errors.New("inventory is full")

func NewInventory() *Inventory {
	items := make(map[int]*Item)
	for i := 0; i < 28; i++ {
		items[i] = blankItem
	}
	return &Inventory{items: items}
}

func validIndex(i int) bool {
	return i >= 0 && i < 28
}

func (i *Inventory) Swap(pos1, pos2 int) error {
	if !validIndex(pos1) || !validIndex(pos2) {
		return ErrBadInventoryIndex
	}
	i.items[pos1], i.items[pos2] = i.items[pos2], i.items[pos1]
	return nil
}

func (i *Inventory) RemoveItem(index int) (*Item, error) {
	if !validIndex(index) {
		return nil, ErrBadInventoryIndex
	}
	removed := i.items[index]
	i.items[index] = blankItem
	return removed, nil
}

func (i *Inventory) UseXItem(index int, amount int) (int, error) {
	if !validIndex(index) {
		return 0, ErrBadInventoryIndex
	}
	toUse := i.items[index]
	if toUse.Quantity == 0 {
		return 0, ErrBadInventoryItemQuantity
	}
	if toUse.Quantity <= amount {
		i.items[index] = blankItem
		return toUse.Quantity, nil
	} else {
		i.items[index] = &Item{ID: toUse.ID, Quantity: toUse.Quantity - amount}
		return amount, nil
	}
}

func (i *Inventory) AddItem(item Item) error {
	if item.ID == blankID {
		return ErrBadInventoryItemID
	}
	// Check for existing stackable item
	if item.Quantity > 0 {
		for _, existingItem := range i.items {
			if existingItem.ID == item.ID {
				existingItem.Quantity += item.Quantity
				return nil
			}
		}
	}
	// Insert in first empty spot
	for index, existingItem := range i.items {
		if existingItem.ID == blankID {
			i.items[index] = &item
			return nil
		}
	}
	return ErrInventoryFull
}
