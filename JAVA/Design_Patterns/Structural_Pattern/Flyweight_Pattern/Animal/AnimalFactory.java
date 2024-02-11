package Design_Patterns.Structural_Pattern.Flyweight_Pattern.Animal;

import java.util.HashMap;

public class AnimalFactory {
    
    private static HashMap<String ,Animal>animalMap = new HashMap<>();

    public Animal getAnimal(String type){

        Animal animal = animalMap.get(type.toLowerCase());
        if (animal != null){
            return animal;
        }
        switch (type.toLowerCase()) {

            case "cat":
                animal = new Cat();
                break;
            
            case "dog":
                animal = new Dog();
                break;
        
            default:
                System.out.println(type + " is not a valid animal type.");
                return null;
        }
        animalMap.put(type.toLowerCase(), animal);
        return  animal;

    }
}
