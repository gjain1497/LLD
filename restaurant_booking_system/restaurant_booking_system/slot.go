package restaurant_booking_system

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
	Date           string
	Time           string
	NumberOfTables int
}
