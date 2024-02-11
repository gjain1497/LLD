package Design_Patterns.Creational_Pattern.Builder_Pattern;

public class Home {
    public String floors;
    public String walls;
    public  String terrace;

    public void displayFeature(){
        System.out.println("Floors : " + this.floors);
        System.out.println("Walls : " + this.walls);
        System.out.println("Terrace : " + this.terrace);
    }
}
