package Design_Patterns.Multithreading;

class MyThread implements Runnable{
    @Override
    public void run() {
        System.out.println("Thread 2 is running");
    }
}
public class runnable_interface{
    public static void main(String[] args) {
        
        MyThread mt = new MyThread();
        
        Thread t = new Thread(mt);
        t.start();

    }
}
