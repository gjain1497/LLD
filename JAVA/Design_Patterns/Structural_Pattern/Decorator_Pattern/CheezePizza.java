package Design_Patterns.Structural_Pattern.Decorator_Pattern;

public class CheezePizza extends ToppingDecorator{
    public  CheezePizza(BasePizza pizza) {
        super(pizza);
    }

    @Override
    public int cost(){
        return this.pizza.cost() + 50;
    }
}
