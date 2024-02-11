package Design_Patterns.Structural_Pattern.Flyweight_Pattern;

import Design_Patterns.Structural_Pattern.Flyweight_Pattern.Animal.Animal;
import Design_Patterns.Structural_Pattern.Flyweight_Pattern.Animal.AnimalFactory;

public class Main {
    public static void main(String[] args) {

        AnimalFactory factory = new AnimalFactory();
        Animal cat1 = factory.getAnimal("Cat");
        cat1.showInfo();

        Animal cat2 = factory.getAnimal("Cat");
        cat2.showInfo();

        System.out.println(cat1);
        System.out.println(cat2);

    }
}