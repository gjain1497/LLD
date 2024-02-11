package Design_Patterns.Creational_Pattern.Singleton_Pattern;

public class Singleton {
    
    private String name, design;
    private static Singleton singleInstance = null;

    private Singleton(){
        this.name = "Singleton";
        this.design = "Creational";
    }
    public static Singleton getInstance(){
        if(singleInstance == null){
            singleInstance = new Singleton();
        }
        return singleInstance;
    }
    public  void showMessage(){
        System.out.println("HashCode of obj " + singleInstance);
        System.out.println(name + " : "+ design);
    }

}
