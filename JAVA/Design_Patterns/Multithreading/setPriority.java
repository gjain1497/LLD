package Design_Patterns.Multithreading;

class X extends Thread{
    @Override
    public void run() {
        for (int i = 0; i < 3; i++) {
            try {
                Thread.sleep(1000);
            } catch (Exception e) {
                // TODO: handle exception
            }
            System.out.println("Hey, I am a "+ Thread.currentThread().getName() +" thread for loop number - " +i);
        }
    }
}
public class setPriority {
    public static void main(String[] args) {
        X t1 = new X();
        X t2 = new X();
        X t3 = new X();

        System.out.println("Before setting priority: " + t1.getPriority()); //


        t1.setName("Child thread 1");
        t2.setName("Child thread 2");
        t3.setName("Child thread 3");

        t1.setPriority(2);
        t2.setPriority(4);
        t3.setPriority(10);

        t1.start();
        t2.start();
        t3.start();

        for (int i = 0; i < 3; i++) {
            try {
                Thread.sleep(1000);
            } catch (Exception e) {
                // TODO: handle exception
            }
            System.out.println("Hey, I am a "+ Thread.currentThread().getName() +" thread for loop number - " +i);
        }
    }
    
}
