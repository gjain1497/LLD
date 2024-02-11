package Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observer;

import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observable.StocksObservable;

public class MobileNotificationObserver implements NotificationObserver{
    
    private  String mobile;
    private String name;

    public MobileNotificationObserver(String name,String mobile) {
        this.name = name;
        this.mobile = mobile;
    }
    @Override
    public  void update(StocksObservable s){
        SendEmail(name,mobile, "Your liked product is in stocks, hurray up!");
        System.out.println();
    }

    public void SendEmail(String name,String to, String message) {

        System.out.println("Sending Message To :"+to);
        try {
            Thread.sleep(1000);
        } catch (Exception e) {
            e.printStackTrace();
        }
        System.out.println("Hey "+name+ ", "+message);
        System.out.println("Message sent to : "+ mobile);

    }
}
