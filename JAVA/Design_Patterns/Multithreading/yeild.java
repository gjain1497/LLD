package Design_Patterns.Multithreading;

class thread1 extends Thread{
    @Override
    public void run(){
        
        for (int i = 0; i < 3; i++) {
            System.out.println("Thread 1 is running " + i);
            Thread.yield();
        }
    }
}
class thread2 extends Thread{
    @Override
    public void run() {
        for (int i = 0; i < 3; i++) {
            System.out.println("Thread 2 is running " + i);
        }    
    }
}

public class yeild {
    public static void main(String[] args) throws InterruptedException {
        
        thread1 t1 = new thread1();
        thread2 t2 = new thread2();

        t1.start();
        Thread.sleep(500);
        t2.start();
        

    }
}
