package LLD.Parking_Lot.models;

import LLD.Parking_Lot.models.vehicle.Vehicle;
import LLD.Parking_Lot.models.vehicle.VehicleType;

public class ParkingSpot {

    private String ID;
    private int floorNo;
    private boolean isAvailable;
    private VehicleType type;
    private Vehicle vehicle;

    ParkingSpot(){

    }
    public void assignVehicle(Vehicle vehicle) {
        this.vehicle = vehicle;
        isAvailable = false;
    }
    public void removeVehicle() {
        this.vehicle = null;
        isAvailable = true;
    }

}
