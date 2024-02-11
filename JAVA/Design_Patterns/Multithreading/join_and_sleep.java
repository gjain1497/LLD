package Design_Patterns.Multithreading;

class D extends Thread{
    @Override
    public void run() {
        
        String name = Thread.currentThread().getName();
        for (int i = 0; i < 3; i++) {
            System.out.println("Thread "+ name + " currently running " + i);
        }
    }
}
public class join_and_sleep {
    public static void main(String[] args) {

        D t1 = new D();
        D t2 = new D();
        D t3 = new D();

        t1.setName("A");
        t2.setName("B");
        t3.setName("C");

        t1.start();
        t2.start();
        t3.start();

        try {
            t1.join();
        } catch (Exception e) {
            // TODO: handle exception
        }

        String name = Thread.currentThread().getName();
        for (int i = 0; i < 3; i++) {
            System.out.println("Thread "+ name + " currently running " + i);
        }
    }
}
