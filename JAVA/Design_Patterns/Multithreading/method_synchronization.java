package Design_Patterns.Multithreading;


class Passenger extends Thread{

    private String name;
    private Bus bus;

    public Passenger(String name,Bus bus) {
        this.name = name;
        this.bus = bus;
    }

    @Override
    public void run() {
        if(this.bus.bookSeat()){
          System.out.println("Seat Available! Booking Successful by "+ this.name);
        } else {
          System.out.println("No Seat Available! Can't Book.");
        }
    }
}
class Bus {

    private int seats = 0;
    Bus(int seats){
        this.seats = seats;
    }
    
    public synchronized boolean bookSeat(){

        if(seats <= 0)
          return false;
        
        System.out.println("Current seats available : "+seats);
        seats -= 1;
        return  true;
    }
}
public class method_synchronization {
    public static void main(String[] args) {

        Bus bus = new Bus(2);

        Passenger p1 = new Passenger("Ram",bus);
        Passenger p2 = new Passenger("Shyam",bus);
        Passenger p3 = new Passenger("Arjun",bus);

        p1.start();
        p2.start();
        p3.start();

    }
}
