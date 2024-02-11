package Design_Patterns.Structural_Pattern.Flyweight_Pattern.Animal;

import Design_Patterns.Structural_Pattern.Flyweight_Pattern.Shared_Class;

public class Cat implements Animal{

    private String type = "Cat";
    public String getType(){
        return  this.type;
    }
    public void showInfo(){
        System.out.println(this.type + " has " + Shared_Class.eyes + " eyes.");
        System.out.println(this.type + " has " + Shared_Class.nose + " nose.");
        System.out.println(this.type + " has " + Shared_Class.legs + " legs.");
        System.out.println(this.type + " has " + Shared_Class.tails + " tails.");
        System.out.println(this.type + " meows");
    }
}
