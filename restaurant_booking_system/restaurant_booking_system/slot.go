package restaurant_booking_system

type SlotStatus int

const (
	maxTables = 4
)

const (
	SLOTAVAILABLE SlotStatus = iota
	SLOTNOTAVAILABLE
)

var SlotStatusString = map[SlotStatus]string{
	SLOTAVAILABLE:    "SLOT_AVAILABLE",
	SLOTNOTAVAILABLE: "SLOT_NOT_AVAILABLE",
}

type Slot struct {
	Date           string
	Time           string
	NumberOfTables int
}

func NewSlot(date string, time string) *Slot {
	return &Slot{
		Date:           date,
		Time:           time,
		NumberOfTables: maxTables,
	}
}
