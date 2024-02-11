package Design_Patterns.Creational_Pattern.Factory_Pattern;

import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Doctor;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Engineer;
import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Profession;

public class Factory {
    
    public Profession getProfession(String profession){
        
        switch (profession.toLowerCase()) {
            case "doctor":
                  return new Doctor();
            case  "engineer":
                return new Engineer();
            case "default":
                System.out.println("Invalid Profession");
        }
        return null;
    }
}
