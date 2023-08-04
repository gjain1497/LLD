package main


ParkingLot{
	Flooor[],
	Entry[],
	Exit[],
	Capacity
	adress Location

	//crud operations
	-add floor
	-update floor
	-remove floor
	-add entry
	-update entry
	-remove entry
	-add exit
	-update exit
	-remove exit
}

Floor{
	Spot[],
	DisplayBoard
	ParkingStrategy

	//crud
	-addSpot
	-removeSpot
	-assignVehicleToSlot
	--deleteVehicleFromSlot
	-addDisplayBoard
	-updateDisplayBoard
	-removeDisplayBoard
}
ParkingStrategy{
	NearToEntrance,
	NearToEntranceAndElevator,
	Default
	find()
}
NearToEntrance{
	find()
}

ParkingTicket{
	ticketNumber
	Vehicle
	parkingSpot
	issuedAt
	payedAt
	payedAmount
	status: ParkingTicketStatus

	//getters and setters

}

ParkingTicketStatus ENUM{
	ACTIVE,
	PAID,
	LOST,
}

Entry{
	ParkingTicket
	ParkingSpot
	id
	findClosestParkingSpot() //because there can be multiple entrances
	updateParkingSpot()
	printTicket()
}

Exit{
	ParkingTicket
	Payment
	ParkingSpot
	scanTicket(): bool
	processPayement(): bool
	updateParkingSpot(): bool
}

Parking Spot{ //DONE
	int id,
	Status: isEmpty or isFilled,
	Vehicle,
	int price,
	Spottype
	Handicaaped,
	Compact,
	Large,
	Motorbike,
	Electric
	ElectricPanel

	-parkVehicle()
	-removeVehicle()
}

TwoWheelerParkingSpot is a ParkingSpot{
	price{10}
	-parkVehicle()
	-removeVehicle()
}

FourWheelerParkingSpot is a ParkingSpot{
	price{20}
	-parkVehicle()
	-removeVehicle()
}



Vehicle {
	vehicle number
	vehicle Vehicletype
}

VehicleType ENUM{
	TRUCK,
	CAR,
	BIKE etc
}

Display Board{
	Parking
}

//addtioanl which I could not think of

ParkingRate{
	hourNumber
	rate
}

Payment{
	creationDate
	amount
	status
	initiateTransaction()
}

ParkingAttendantPortal{

}

CustomerInfoPortal{

}
ElectricPanel{

}