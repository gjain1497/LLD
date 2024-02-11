package Design_Patterns.Creational_Pattern.Prototype_Pattern.Profession;

public class Engineer extends Profession{

    public Engineer(int id,String name){
        super(id, name);
    }
    public void displayProfession(){
        System.out.println("I am an engineer.");
    }
}
