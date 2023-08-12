package main

type SlotStatus int

const (
	SLOTAVAILABLE SlotStatus = iota
	SLOTNOTAVAILABLE
)

var SlotStatusString = map[SlotStatus]string{
	SLOTAVAILABLE:    "SLOTAVAILABLE",
	SLOTNOTAVAILABLE: "SLOTNOTAVAILABLE",
}

type Slot struct {
	Date           int
	Time           int
	Status         SlotStatus
	NumberOfPeople int
}
