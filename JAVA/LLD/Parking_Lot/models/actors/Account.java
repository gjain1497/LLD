package LLD.Parking_Lot.models.actors;

import LLD.Parking_Lot.models.Person;

public abstract class Account {

    private String ID;
    protected String name,password;
    private Person person;
    protected AccountStatus  status;

    // If the method to resetPassword is different for different class extending this abstract class
    // then members which can be modified can be kept protected but if the logic is same then we can 
    // keep the method protected
    public abstract void resetPassword(String password);

    // Other approach
    // protected void resetPassword(String password){
    //     thios.password = password;
    // }
    
}
