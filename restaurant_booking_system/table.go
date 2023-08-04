package main

type TableType int

const (
	TwoSeater TableType = iota
	FourSeater
	Family
)

type Table struct {
	TableId    int
	PersonList []*Person
	TableType  TableType
	Slot       *Slot
}
