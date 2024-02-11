package LLD.Parking_Lot.models.vehicle;

import LLD.Parking_Lot.models.ParkingTicket;

public class Car extends Vehicle{

    public Car(String registrationNumber) {
        super(VehicleType.CAR,registrationNumber);
    }
    public  void assignTicket(ParkingTicket ticket){
        super.assignTicket(ticket);
    }
}
