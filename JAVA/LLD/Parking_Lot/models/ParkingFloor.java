package LLD.Parking_Lot.models;

import LLD.Parking_Lot.models.vehicle.VehicleType;
import java.util.*;

public class ParkingFloor {

    private int floorNo;

    // Stores count of free spots for each vehicle type per floor in the parking lot
    HashMap<Integer,HashMap<VehicleType,Integer>> floorParkingSpots;
    private ParkingDisplayBoard displayBoard;

    public ParkingFloor(int number,ParkingDisplayBoard displayBoard){
        this.floorNo = number;
        this.displayBoard = displayBoard;
    }

    public void addParkingSpot(int floor,ParkingSpot  spot) throws Exception{
        
    }
    public void assignVehicleToSpot(int floor,VehicleType type){
        
    }
    public void updateDisplayBoard(ParkingSpot spot){

    }
    public void freeSpot(ParkingSpot spot){
        spot.removeVehicle();
    }

}
