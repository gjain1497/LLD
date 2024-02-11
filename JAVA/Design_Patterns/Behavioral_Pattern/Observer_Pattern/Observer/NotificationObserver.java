package Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observer;

import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observable.StocksObservable;

public interface NotificationObserver {

    public  void update(StocksObservable s);
} 