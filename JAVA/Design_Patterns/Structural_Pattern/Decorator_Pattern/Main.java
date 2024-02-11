package Design_Patterns.Structural_Pattern.Decorator_Pattern;

public class Main {

    public static void main(String[] args){

        Farmhouse f1 = new Farmhouse();
        System.out.println("FarmHouse Cost: " + f1.cost());
        
        CheezePizza ch = new CheezePizza(f1);
        System.out.println("CheezePizza Cost: " + ch.cost());

    }
}