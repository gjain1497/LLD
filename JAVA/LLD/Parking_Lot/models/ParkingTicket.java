package LLD.Parking_Lot.models;

import java.util.*;

import LLD.Parking_Lot.models.gate.EntryGate;
import LLD.Parking_Lot.models.gate.ExitGate;
import LLD.Parking_Lot.models.vehicle.Vehicle;

public class ParkingTicket {
    private String ticketNO;
    private Date entryTimestamp;
    private Date exitTimestamp;
    private Vehicle vehicle;
    private  double amount;
    private EntryGate entryGate;
    private ExitGate exitGate;

    public String getTicketNumber(){
        return this.ticketNO;
    }
    public void saveInDB(){
        System.out.println("-- Saving to DB --");
        System.out.println("-- Saved to DB --");
    }
}
