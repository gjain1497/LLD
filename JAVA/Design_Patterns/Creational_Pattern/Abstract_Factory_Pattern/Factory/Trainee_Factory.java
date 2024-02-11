package Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Factory;

import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Profession;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Trainee_Doctor;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Trainee_Engineer;

public class Trainee_Factory extends Factory {

    public Profession getProfession(String profession){
        switch (profession.toLowerCase()) {
            case "doctor":
                  return new Trainee_Doctor();
            case  "engineer":
                return new Trainee_Engineer();
            case "default":
                System.out.println("Invalid Profession");
        }
        return null;
    }
}
