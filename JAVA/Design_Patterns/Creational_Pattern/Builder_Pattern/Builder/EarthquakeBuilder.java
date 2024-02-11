package Design_Patterns.Creational_Pattern.Builder_Pattern.Builder;

import Design_Patterns.Creational_Pattern.Builder_Pattern.Home;

public class EarthquakeBuilder implements Builder{
    
    private Home home = new Home();  

    public void buildFloors() {
        this.home.floors = "Wooden floors";
    }

    public void buildWalls() {
        this.home.walls = "Wooden walls";      
    }

    public void buildTerrace() {
        this.home.terrace = "Wooden terrace";
    }

    public Home getComplextHomeObject(){
        return  this.home;
    }
}
