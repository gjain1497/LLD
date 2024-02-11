package Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observable;

import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observer.NotificationObserver;

public interface StocksObservable {

    public void add(NotificationObserver observer);
    public void remove(NotificationObserver observer);
    public void notifyObservers();
    public void setStocksCount(int count);

}
