package Design_Patterns.Creational_Pattern.Prototype_Pattern;

import Design_Patterns.Creational_Pattern.Prototype_Pattern.Profession.Profession;

public class Main {
    public static void main(String[] args) {

        Profession_Cache  cache = new Profession_Cache();
        cache.loadCache();

        int docID = 828732;

        Profession doc1 = cache.getClonedProfession(docID);
        System.out.println("Doc1: " + doc1);

        Profession doc2 = cache.getClonedProfession(docID);
        System.out.println("Doc2: " + doc2);

    }
}
