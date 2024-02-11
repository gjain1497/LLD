package Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern;

import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Profession;

public class Main {
    public static void main(String[] args) {
        Abstract_Factory factory = new Abstract_Factory();
        
        Profession eng = factory.getFactory("trainee").getProfession("engineer");
        Profession doc = factory.getFactory("profession").getProfession("doctor");

        eng.displayProfession();
        doc.displayProfession();

    }
}
