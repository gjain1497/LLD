package Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern;

import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Factory.Factory;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Factory.Profession_Factory;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Factory.Trainee_Factory;

public class Abstract_Factory {
    
    public Factory  getFactory(String choice) {

        switch (choice.toLowerCase()) {
            case "trainee":
                  return new Trainee_Factory();
            case  "profession":
                return new Profession_Factory();
            case "default":
                System.out.println("Invalid Profession");
        }
        return null;
    }
}
