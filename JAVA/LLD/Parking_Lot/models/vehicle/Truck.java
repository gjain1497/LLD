package LLD.Parking_Lot.models.vehicle;

import LLD.Parking_Lot.models.ParkingTicket;

public class Truck extends Vehicle{
    
    public Truck(String registrationNumber) {
        super(VehicleType.TRUCK,registrationNumber);
    }
    public  void assignTicket(ParkingTicket ticket){
        super.assignTicket(ticket);
    }
}
