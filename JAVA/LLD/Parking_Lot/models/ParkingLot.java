package LLD.Parking_Lot.models;

import java.util.*;
import LLD.Parking_Lot.models.gate.*;
import LLD.Parking_Lot.models.vehicle.*;

public class ParkingLot {
    
    private String name;
    private Address Address;
    private ParkingRate parkingRate;

    ArrayList<EntryGate>entryPanels;
    ArrayList<ExitGate>exitPanels;

    HashMap<String,ParkingFloor>parkingFloors;
    HashMap<String,ParkingTicket>activeTickets;
    HashMap<VehicleType,Integer>currFreeSpots;

    private static ParkingLot parkingLot = null;

    private ParkingLot(){

        this.name = "Parking Lot";
        this.Address = new Address();
        this.parkingRate = new ParkingRate();
    }

    public static ParkingLot getInstance() { 

        if(parkingLot == null){
            parkingLot = new ParkingLot();
        }
        return parkingLot;
    }

    public static void main(String[] args) {

        Address a = new Address("12E MacBrie Road","Jersey","New York","US",99032);
        ParkingLot myParkingLot = ParkingLot.getInstance();

        Car c1 = new Car("BHI657889");
        Car c2 = new Car("MJHI76659");

        Truck t1 = new Truck("USJS657889");
        Truck t2 = new Truck("JKMI76659");

        MotorBike m1 = new MotorBike("BHI657889");
        MotorBike m2 = new MotorBike("MJHI76659");

        HashMap<VehicleType, Integer> floor1Data = new HashMap<VehicleType, Integer>() {{
            put(VehicleType.CAR, 2);
            put(VehicleType.TRUCK, 1);
            put(VehicleType.MOTORBIKE, 3);
        }};

        HashMap<VehicleType, Integer> floor2Data = new HashMap<VehicleType, Integer>() {{
            put(VehicleType.CAR, 0);
            put(VehicleType.TRUCK, 1);
            put(VehicleType.MOTORBIKE, 1);
        }};

        // ParkingDisplayBoard displayBoard = new ParkingDisplayBoard(null, null);
        ParkingFloor parkingFloor = new ParkingFloor(2, null);

    }

    public synchronized ParkingTicket getParkingTicket(Vehicle v) throws Exception {

        if(this.isFull(v.getType())) {
            throw new Exception("Parking space full");
        }
        
        ParkingTicket t = new ParkingTicket();
        v.assignTicket(t);
        t.saveInDB();

        this.decrementFreeSpots(v.getType());
        this.activeTickets.put(t.getTicketNumber(), t);
        return t;

    }


    public boolean isFull(VehicleType type){
        return currFreeSpots.get(type) <= 0;
    }

    public void decrementFreeSpots(VehicleType type){

        currFreeSpots.put(type,currFreeSpots.get(type)-1);
        return;
    }


    public void addParkingFloor(){
        // Add parking floor
    }

    public void addEntryPanel(){

    }
    public void addExitPanel(){

    }

}
