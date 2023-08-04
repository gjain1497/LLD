package main

type Store struct {
	StoreID             int
	InventoryManagement *VehicleInventoryManagement
	StoreLocation       Location
	Reservations        []*Reservation
}

func (s *Store) GetVehicles(vehicleType VehicleType) []*Vehicle {
	vehicles := s.InventoryManagement.GetVehicles(vehicleType)
	return vehicles
}

func (s *Store) SetVehicles(vehicles []*Vehicle) {
	s.InventoryManagement = NewVehicleInventoryManagement(vehicles)
}

func (s *Store) CreateReservation(vehicle *Vehicle, user *User) *Reservation {
	reservation := &Reservation{}
	reservation.createReservation(user, vehicle)
	s.Reservations = append(s.Reservations, reservation)
	return reservation
}

func (s *Store) UpdateReservation(reservationID int) bool {
	// Find the reservation in the list
	ReservationObj := &Reservation{}
	for _, reservation := range s.Reservations {
		if reservation.ReservationId == reservationID {
			ReservationObj = reservation
			break
		}
	}
	ReservationObj.updateReservation()
	return true
}

func (s *Store) CompleteReservation(reservationID int) bool {
	// Find the reservation in the list and call the completeReservation method.
	ReservationObj := &Reservation{}
	for _, reservation := range s.Reservations {
		if reservation.ReservationId == reservationID {
			ReservationObj = reservation
			break
		}
	}
	ReservationObj.completeReservation()
	return true
}

func (s *Store) CancelReservation(reservationID int) bool {
	// Find the reservation in the list and call the completeReservation method.
	ReservationObj := &Reservation{}
	for _, reservation := range s.Reservations {
		if reservation.ReservationId == reservationID {
			ReservationObj = reservation
			break
		}
	}
	ReservationObj.cancelReservation()
	return true
}
