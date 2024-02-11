package Design_Patterns.Behavioral_Pattern.Observer_Pattern;

import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observable.*;
import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observer.*;


public class Main {
    public static void main(String[] args) {
        
        // Observable
        StocksObservable iPhone13 = new IphoneStocksObservable("US","13");
        StocksObservable iPhone14Pro = new IphoneStocksObservable("US","14 Pro");

        // Observers
        NotificationObserver  JohnDoe = new EmailNotificationObserver("John Doe", "john@gmail.com");
        NotificationObserver  JaneSmith = new EmailNotificationObserver("Jane Smith", "jane@yahoo.com");
        NotificationObserver   MrRobot = new MobileNotificationObserver("Mr Robot", "+1-202-555-0173");

        // Adding observers to observable
        iPhone13.add(MrRobot);
        iPhone13.add(JohnDoe);
        iPhone14Pro.add(JaneSmith);

        // Adding items in stocks
        iPhone13.setStocksCount(10);
        iPhone14Pro.setStocksCount(5);

    }
}
