package Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observer;

import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observable.StocksObservable;

public class EmailNotificationObserver implements NotificationObserver{
    
    private  String emailID;
    private String name;
    public EmailNotificationObserver(String name,String email) {
        this.name = name;
        this.emailID = email;
    }
    @Override
    public  void update(StocksObservable s){
        SendEmail(emailID, "Your product is in stocks, hurray up!");
        System.out.println();
    }

    public void SendEmail(String to, String message) {
        System.out.println("Sending Email To :"+to);
        try {
            Thread.sleep(1000);
        } catch (Exception e) {
            e.printStackTrace();
        }
        System.out.println("Hey "+name+ " , "+message);
        System.out.println("Message sent to : "+ emailID);
    }
}
