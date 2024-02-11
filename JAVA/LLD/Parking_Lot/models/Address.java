package LLD.Parking_Lot.models;

public class Address {

    private String  streetName;
    private String city,  state, country;
    private int zipCode;

    public Address(){
        
    }
    public Address(String streetName, String city, String state, String country, int zipCode) {
        this.streetName = streetName;
        this.city = city;
        this.state = state;
        this.country = country;
        this.zipCode = zipCode;
    }
    
}