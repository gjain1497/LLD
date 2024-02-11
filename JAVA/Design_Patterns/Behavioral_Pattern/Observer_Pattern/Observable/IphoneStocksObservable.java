package Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observable;

import java.util.ArrayList;
import java.util.List;

import Design_Patterns.Behavioral_Pattern.Observer_Pattern.Observer.NotificationObserver;

public class IphoneStocksObservable implements StocksObservable{
    
    private List<NotificationObserver> observers = new ArrayList<>();
    private int stocksCount = 0;
    private String product = "IPhone";
    private String  country;
    private String  model;

    public IphoneStocksObservable(String country, String model) {
        this.country  = country;
        this.model = model;
    }
    public String getModel(){
        return this.model;
    }

    @Override
    public void add(NotificationObserver obj) {
        observers.add(obj);
    }

    @Override
    public void remove(NotificationObserver obj) {
        observers.remove(obj);
    }

    @Override
    public void notifyObservers(){
        for(NotificationObserver obj : observers){
            obj.update(this);
        }
    }

    public void setStocksCount(int count){
        if(stocksCount == 0) 
            notifyObservers();
        
        stocksCount = count;
    }
}
