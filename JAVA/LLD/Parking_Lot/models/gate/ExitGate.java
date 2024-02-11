package LLD.Parking_Lot.models.gate;

import LLD.Parking_Lot.models.ParkingTicket;

public class ExitGate extends Gate{
    ExitGate(int id){
        super(id);
    }

    public void open(){

    }
    public void close(){

    } 

    public void validateTicket(ParkingTicket ticket){
        // do validation
    }

}
