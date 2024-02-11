package Design_Patterns.Multithreading;

class T extends Thread{
    @Override
    public void run() {
        
        // Prints 5 time as throws interrupted exception  if thread is interrupted while sleeping
        // Main thread inturrupts for single time so caught exception single time
        for (int i = 0; i < 5; i++) {
            System.out.println("Child thread running --- ");
            try {
                Thread.sleep(1000);
            } catch (Exception e) {
                System.out.println("Exception caught");
            }
        }

        // Prints single time 
        // try {
        //     for(int i=0;i<5;i++){
        //         System.out.println("Child thread running --- ");
        //         Thread.sleep(1000);
        //     }
        // } catch (Exception e) {
        //     System.out.println("Exception caught");
        // }
    }
}
public class interrupt {
    public static void main(String[] args) {
        
        T t = new T();
        t.start();
        t.interrupt();
        System.out.println("Main thread is running --- ");
    }
}
