package Design_Patterns.Creational_Pattern.Builder_Pattern.Builder;

import Design_Patterns.Creational_Pattern.Builder_Pattern.Home;

public interface Builder {

    public void buildFloors();
    public void buildWalls();
    public void buildTerrace();
    public Home getComplextHomeObject();

} 