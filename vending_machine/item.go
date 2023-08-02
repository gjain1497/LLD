package main

type Item struct {
	Type  ItemType
	Price int
}

func (i Item) GetType() ItemType {
	return i.Type
}

func (i Item) SetType(t ItemType) {
	i.Type = t
}

func (i Item) GetPrice() int {
	return i.Price
}

func (i Item) SetPrice(price int) {
	i.Price = price
}
