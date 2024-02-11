package Design_Patterns.Creational_Pattern.Prototype_Pattern.Profession;

public abstract class Profession implements Cloneable {

    private int id;
    private String name;
    
    Profession(int id,String name){
        this.id = id;
        this.name = name;
    }

    public int getId(){
        return this.id;
    }
    public String getName(){
        return this.name;
    }

    public Object clone() {

        Object obj = null;
        try {
            obj = super.clone();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return obj;
    }
    abstract void displayProfession();
    
} 
