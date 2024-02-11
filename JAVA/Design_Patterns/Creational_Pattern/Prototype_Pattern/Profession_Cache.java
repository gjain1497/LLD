package Design_Patterns.Creational_Pattern.Prototype_Pattern;

import java.util.HashMap;

import Design_Patterns.Creational_Pattern.Prototype_Pattern.Profession.*;

public class Profession_Cache {
    
    private HashMap<Integer, Profession> cache = new HashMap<>();
    public void loadCache(){

        Profession doc = new Doctor(828732, "John Doe");
        cache.put(doc.getId(), doc);
        Profession eng = new Engineer(383881, "Jane Smith");
        cache.put(eng.getId(), eng);
    }
    public Profession getClonedProfession (int professionId){
        if (!cache.containsKey(professionId)) {
            System.out.println("No such profession in the cache!");
            return null;
        } 
        return  (Profession) cache.get(professionId).clone();
    }

}
