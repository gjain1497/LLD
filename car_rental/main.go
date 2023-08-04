package main

func main() {
	users := addUsers()

	car := &Car{
		Vehicle: Vehicle{
			VehicleID:   4158,
			VehicleType: Caar,
		},
		NumberOfDoors: 4,
		IsConvertible: false,
	}

	//vehicles := addVehicles()
	car.Drive()
	stores := addStores(vehicles)

	rentalSystem := NewVehicleRentalSystem(stores, users)

	// 0. User comes
	user := users[0]

	// 1. User searches store based on location
	location := NewLocation(403012, "Bangalore", "Karnataka", "India")
	store := rentalSystem.getStore(location)

	// 2. Get all vehicles the user is interested in (based on different filters)
	storeVehicles := store.GetVehicles(Caar)

	// 3. Reserving the particular vehicle from vehicle list
	reservation := store.CreateReservation(storeVehicles[0], user)

	// 4. Generate the bill
	bill := NewBill(reservation)

	// 5. Make payment
	payment := &Payment{}
	payment.payBill(bill)

	//6. Trip in progress
	store.UpdateReservation(reservation.ReservationId)

	cancel := true
	//7 if user wants to cancel
	if cancel {
		store.CancelReservation(reservation.ReservationId)
	}

	// 8. Trip completed, submit the vehicle and close the reservation
	store.CompleteReservation(reservation.ReservationId)
}


func addVehicles() []*Vehicle {
	vehicles := []*Vehicle{
		&Car{
			Vehicle: Vehicle{
				VehicleID:   4158,
				VehicleType: Caar,
			},
			NumberOfDoors: 4
			IsConvertible: false
		}
		&Bike{
			Vehicle: Vehicle{
				VehicleID:   6993,
				VehicleType: Motorcycle,
			},
			HasStorageBox: true,
		}
	}
	return vehicles
}

func addUsers() []*User {
	users := []*User{
		{
			UserID:         1,
			UserName:       "Girish Jain",
			DrivingLicense: 125,
		},
	}
	return users
}

func addStores(vehicles []*Vehicle) []*Store {
	store := &Store{StoreID: 1}
	store.SetVehicles(vehicles)

	stores := []*Store{store}

	return stores
}
