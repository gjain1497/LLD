package main

import "errors"

type Inventory struct {
	Itemshelf []ItemShelf
}

func NewInventory(itemCount int) Inventory {
	inventory := &Inventory{}
	inventory.Itemshelf = make([]ItemShelf, itemCount)
	return inventory
}

func (inv Inventory) GetItemShelf() []ItemShelf {
	return inv.Itemshelf
}

func (inv Inventory) SetItemShelf(itemShelf []ItemShelf) {
	inv.Itemshelf = itemShelf
}

func (inv Inventory) InitialEmptyInventory() {
	startCode := 101
	for i := 0; i < len(inv.Itemshelf); i++ {
		space := &ItemShelf{}
		space.SetCode(startCode)
		space.SetSoldOut(true)
		inv.Itemshelf[i] = space
		startCode++
	}
}

func (inv Inventory) AddItem(item Item, codeNumber int) error {
	for _, itemShelf := range inv.Itemshelf {
		if itemShelf.GetCode() == codeNumber {
			if itemShelf.IsSoldOut() {
				itemShelf.SetItem(item)
				itemShelf.SetSoldOut(false)
			} else {
				return errors.New("already item is present, you cannot add item here")
			}
		}
	}
	return nil
}
func (inv Inventory) GetItem(codeNumber int) (Item, error) {
	var result Item
	for _, itemShelf := range inv.Itemshelf {
		if itemShelf.GetCode() == codeNumber {
			if itemShelf.IsSoldOut() {
				return result, errors.New("item already sold out")
			}
			return itemShelf.GetItem(), nil
		}
	}
	return result, errors.New("invalid code")
}

func (inv Inventory) UpdateSoldOutItem(codeNumber int) {
	for _, itemShelf := range inv.Itemshelf {
		if itemShelf.GetCode() == codeNumber {
			itemShelf.SetSoldOut(true)
		}
	}
}
