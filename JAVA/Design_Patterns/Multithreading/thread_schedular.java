package Design_Patterns.Multithreading;

class C extends Thread {
    @Override
    public void run() {
        String name = Thread.currentThread().getName();
        for (int i = 0; i < 3; i++) {
            System.out.println("Curr thread: " + name + ", Value: " + i);
        }
    }
}

public class thread_schedular {
    public static void main(String[] args) {

        C t1 = new C();
        C t2 = new C();
        C t3 = new C();

        t1.setName("Thread 1");
        t2.setName("Thread 2");
        t3.setName("Thread 3");

        t1.start();
        t2.start();
        t3.start();
    }
}

