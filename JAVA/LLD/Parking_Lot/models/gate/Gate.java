package LLD.Parking_Lot.models.gate;

public abstract class Gate {
    
    private int id;
    Gate(int id){
        this.id  = id;
    }
    abstract public void open();
    abstract public void close();

}
