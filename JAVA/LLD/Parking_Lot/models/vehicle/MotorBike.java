package LLD.Parking_Lot.models.vehicle;

import LLD.Parking_Lot.models.ParkingTicket;

public class MotorBike extends Vehicle{
    
    public MotorBike(String registrationNumber) {
        super(VehicleType.MOTORBIKE,registrationNumber);
    }
    public  void assignTicket(ParkingTicket ticket){
        super.assignTicket(ticket);
    }
    
}
