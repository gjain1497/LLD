package LLD.Parking_Lot.models;

import java.util.HashMap;
import LLD.Parking_Lot.models.vehicle.VehicleType;

public class ParkingDisplayBoard {
    
    private String ID;

    // Can use singleton pattern
    ParkingDisplayBoard(String ID){
        this.ID = ID;
    }

    public void display(HashMap<Integer,HashMap<VehicleType,Integer>>floorParkingSpots){
        for ( Integer floorNo : floorParkingSpots.keySet()) {
            System.out.println("-- Available slots on floor : "+ floorNo + " --") ;
            for(VehicleType type :  VehicleType.values() ){
                System.out.println( type + "  : "+floorParkingSpots.get(floorNo).get(type));
            }
        }
    }
}
