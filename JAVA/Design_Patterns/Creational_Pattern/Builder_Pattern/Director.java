package Design_Patterns.Creational_Pattern.Builder_Pattern;

import Design_Patterns.Creational_Pattern.Builder_Pattern.Builder.Builder;

public class Director {
    private Builder builder;

    public Director(Builder builder) {
        this.builder = builder;
    }

    public Home getHome() {
        manageHomeConstruction();
        return this.builder.getComplextHomeObject();
    }

    public void manageHomeConstruction(){
        this.builder.buildFloors();
        this.builder.buildWalls();
        this.builder.buildTerrace();
    }
}
