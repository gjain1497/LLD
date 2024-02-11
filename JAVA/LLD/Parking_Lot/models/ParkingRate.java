package LLD.Parking_Lot.models;

public class ParkingRate {
    
    private double hours , rate;
    public double calculate(){
        return this.hours * this.rate ;
    }
}
