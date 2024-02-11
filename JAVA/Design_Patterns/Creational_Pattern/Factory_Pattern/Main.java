package Design_Patterns.Creational_Pattern.Factory_Pattern;

import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Factory.Profession_Factory;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Profession;

public class Main {
    public static void main(String[] args) {
        // Factory pattern is used to create objects without specifying the exact class of object that will be created
        Profession_Factory factory = new Profession_Factory();

        Profession doc = factory.getProfession("doctor");
        Profession eng = factory.getProfession("engineer");

        doc.displayProfession();
        eng.displayProfession();
    }
}
