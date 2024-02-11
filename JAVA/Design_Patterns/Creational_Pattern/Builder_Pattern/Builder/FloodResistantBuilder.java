package Design_Patterns.Creational_Pattern.Builder_Pattern.Builder;

import Design_Patterns.Creational_Pattern.Builder_Pattern.Home;

public class FloodResistantBuilder implements Builder{

    private Home home = new Home();  
    
    public void buildFloors() {
        this.home.floors = "Water Resitant floors";
    }

    public void buildWalls() {
        this.home.walls = "Water Resitant walls";      
    }

    public void buildTerrace() {
        this.home.terrace = "Water Resitant terrace";
    }

    public Home getComplextHomeObject(){
        return  this.home;
    }
}
