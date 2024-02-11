package LLD.Parking_Lot.models.actors;

import LLD.Parking_Lot.models.ParkingDisplayBoard;
import LLD.Parking_Lot.models.ParkingSpot;
import LLD.Parking_Lot.models.gate.*;

public class Admin extends Account {
    
  public void addParkingSpot(ParkingSpot spot){

  }

  public void addDisplayBoard(ParkingDisplayBoard displayBoard){

  }

  public void addEntrance(EntryGate entrance){

  }

  public void addExit(ExitGate exit){

  }

  public void resetPassword(String password) {
     this.password = password;
  }
}
