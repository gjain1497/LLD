package main

type BookingStatus int

const (
	AVAILABLE BookingStatus = iota
	BOOKED
)

var bookingStatusStrings = map[BookingStatus]string{
	AVAILABLE: "AVAILABLE",
	BOOKED:    "BOOKED",
}

type Booking struct {
	BookingId int
	Status    BookingStatus
}
