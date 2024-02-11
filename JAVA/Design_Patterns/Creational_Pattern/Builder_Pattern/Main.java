package Design_Patterns.Creational_Pattern.Builder_Pattern;

import Design_Patterns.Creational_Pattern.Builder_Pattern.Builder.EarthquakeBuilder;

public class Main {
    public static void main(String[] args) {

        EarthquakeBuilder builder = new EarthquakeBuilder();
        
        Director dir  = new Director(builder);

        Home earthQuakHome = dir.getHome();
        System.out.println(earthQuakHome.floors);

    }
}
