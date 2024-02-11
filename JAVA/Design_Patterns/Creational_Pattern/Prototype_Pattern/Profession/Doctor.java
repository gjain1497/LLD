package Design_Patterns.Creational_Pattern.Prototype_Pattern.Profession;

public class Doctor extends Profession{

    public Doctor(int id,String  name) {
        super(id,name);
    }
    public void displayProfession(){
        System.out.println("I am a doctor and I heal the sick");
    }
}
