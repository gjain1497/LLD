package Design_Patterns.Creational_Pattern.Singleton_Pattern;

public class Main {
    public static void main(String[] args) {
        
        Singleton obj1 =  Singleton.getInstance();
        Singleton obj2 =  Singleton.getInstance();

        obj1.showMessage();
        obj2.showMessage();
    }
}
