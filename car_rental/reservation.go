package main

import (
	"time"
)

type ReservationType int

const (
	DAILY ReservationType = iota
	HOURLY
	// Add other reservation types here if needed
)

type ReservationStatus int

const (
	SCHEDULED ReservationStatus = iota
	INPROGRESS
	COMPLETED
	CANCELLED
	// Add other reservation statuses here if needed
)

type Reservation struct {
	reservationId     int
	user              *User
	vehicle           *Vehicle
	bookingDate       time.Time
	dateBookedFrom    time.Time
	dateBookedTo      time.Time
	fromTimeStamp     int64
	toTimeStamp       int64
	pickUpLocation    *Location
	dropLocation      *Location
	reservationType   ReservationType
	reservationStatus ReservationStatus
	location          *Location
}

func (r *Reservation) createReserve(user *User, vehicle *Vehicle) int {
	// generate new id
	r.reservationId = 12232
	r.user = user
	r.vehicle = vehicle
	r.reservationType = DAILY
	r.reservationStatus = SCHEDULED

	return r.reservationId
}

// CRUD operations
