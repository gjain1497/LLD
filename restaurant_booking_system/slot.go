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
	FromTime int
	ToTime   int
	Status   SlotStatus
}
