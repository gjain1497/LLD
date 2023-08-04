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

//type ReservationRentalInsurance int
//
//const (
//	FULLYINSURED ReservationRentalInsurance = iota
//	PARTLYINSURED
//	// Add other reservation statuses here if needed
//)

type Reservation struct {
	ReservationId              int
	User                       *User
	Vehicle                    *Vehicle
	BookingDateTime            time.Time //konsi dtae ko book kia hai
	ScheduledDateTime          time.Time //konsi dtae ko book kia hai
	CompletionDateTime         time.Time //konsi dtae ko book kia hai
	CancellationDateTime       time.Time //konsi dtae ko book kia hai
	FromTimeStamp              int       //kb se lekar kb tk book kia hai
	ToTimeStamp                int       //kb se lekar kb tk book kia hai
	PickUpLocation             *Location
	DropLocation               *Location
	ReservationType            ReservationType
	ReservationStatus          ReservationStatus
	ReservationRentalInsurance Insurance
	Location                   *Location
}

func NewPersonalInsurance() *PersonalInsurance {
	return &PersonalInsurance{
		Insurance: Insurance{
			Provider:   "SecureRide",
			Coverage:   "Comprehensive",
			CostPerDay: 15.0,
		},
		MedicalCoverage: 5000.0,
	}
}

// CRUD operations

func (r *Reservation) createReservation(user *User, vehicle *Vehicle) {
	// generate new id
	r.ReservationId = 12232
	r.User = user
	r.Vehicle = vehicle
	r.ReservationType = DAILY
	r.ReservationStatus = SCHEDULED
	r.ReservationRentalInsurance = NewPersonalInsurance()
	r.BookingDateTime = time.Now()
	r.FromTimeStamp = 20230804
	r.ToTimeStamp = 20230808
	r.Location = &Location{
		Pincode: 403012,
		City:    "Bangalore",
		State:   "Karnataka",
		Country: "India",
	}
}

func (r *Reservation) updateReservation() {
	r.ReservationStatus = INPROGRESS
	r.ScheduledDateTime = time.Now()
}

func (r *Reservation) cancelReservation() {
	r.ReservationStatus = CANCELLED
	r.CancellationDateTime = time.Now()
}

func (r *Reservation) completeReservation() {
	r.User = nil
	r.Vehicle = nil
	r.ReservationStatus = COMPLETED
	r.CompletionDateTime = time.Now()
}
