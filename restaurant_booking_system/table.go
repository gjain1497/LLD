package main

type TableType int

const (
	TwoSeater TableType = iota
	FourSeater
	Family
)

var TableTypeStrings = map[TableType]string{
	TwoSeater:  "TwoSeater",
	FourSeater: "FourSeater",
	Family:     "Family",
}

type Table struct {
	TableId    int
	PersonList []*Person
	TableType  TableType
	Booking    *Booking
	Slots      []*Slot
}
