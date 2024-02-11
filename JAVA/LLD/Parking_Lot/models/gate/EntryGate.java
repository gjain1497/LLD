package LLD.Parking_Lot.models.gate;

import LLD.Parking_Lot.models.ParkingTicket;
import LLD.Parking_Lot.models.vehicle.Vehicle;

public class EntryGate extends Gate {
    
    EntryGate(int id){
        super(id);
    }

    public void open(){

    }
    public void close(){

    } 

    public ParkingTicket getParkingTicket(Vehicle v){

        ParkingTicket t =  new ParkingTicket();
        return t;
    }
}
