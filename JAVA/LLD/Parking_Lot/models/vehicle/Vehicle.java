package LLD.Parking_Lot.models.vehicle;

import LLD.Parking_Lot.models.ParkingTicket;


public abstract class Vehicle {

    private String registrationNumber;
    private final VehicleType type;
    private ParkingTicket ticket;

    public Vehicle(VehicleType type,String number) {
        this.type = type;
        this.registrationNumber = number;
    }

    public void assignTicket(ParkingTicket ticket) {
        this.ticket = ticket;
    }

    public VehicleType getType(){
        return this.type;
    }
    
}
