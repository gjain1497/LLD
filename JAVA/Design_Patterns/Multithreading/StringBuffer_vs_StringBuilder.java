package Design_Patterns.Multithreading;

public class StringBuffer_vs_StringBuilder {
    public static void main(String[] args) {

        StringBuilder sb = new StringBuilder("Java");
        StringBuffer sbf = new StringBuffer("Java");

        Thread therad1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                sb.append("#");
                sbf.append("#");
            }
        });
        Thread therad2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                sb.append("#");
                sbf.append("#");
            }
        });

        therad1.start();
        therad2.start();

        try {
            therad1.join();
            therad2.join();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }

        System.out.println("After thread execution stringBuilder's length : "+sb.length());
        System.out.println("After thread execution sringBuffer's length : "+sbf.length());

    }
}
