package Design_Patterns.Multithreading;

class A extends Thread{

    @Override
    public void run() {
        for (int i = 0; i < 5; i++) {
            System.out.println("Hey, Its MyThread "+ i);
            try {
                Thread.sleep(1000);
            } catch (Exception e) {
                System.out.println("Error occurred");
            }
        }
    }

    public void fun() {
        for (int i = 0; i < 5; i++) {
            System.out.println("Hey, Its MyThread "+ i);
            try {
                Thread.sleep(1000);
            } catch (Exception e) {
                System.out.println("Error occurred");
            }
        }
    }
}
public class thread_class {
    public static void main(String[] args) throws  InterruptedException {
        
         A t = new A();
         t.start();
        //  t.fun();

         for (int i = 0; i < 5; i++) {
            System.out.println("Hey ,  Its Main thread " + i);
            Thread.sleep(1000);
         }
    }
}
