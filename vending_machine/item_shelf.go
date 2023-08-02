package main

type ItemShelf struct {
	item    Item
	code    int
	soldOut bool
}

// Getter methods

func (is ItemShelf) GetItem() Item {
	return is.item
}

func (is ItemShelf) GetCode() int {
	return is.code
}

func (is ItemShelf) IsSoldOut() bool {
	return is.soldOut
}

// Setter methods

func (is ItemShelf) SetItem(item Item) {
	is.item = item
}

func (is ItemShelf) SetCode(code int) {
	is.code = code
}

func (is ItemShelf) SetSoldOut(soldOut bool) {
	is.soldOut = soldOut
}
