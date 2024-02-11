package Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Factory;

import Design_Patterns.Creational_Pattern.Abstract_Factory_Pattern.Profession.Profession;

public abstract class Factory {
    public abstract Profession getProfession(String  profession);
}
