package Design_Patterns.Structural_Pattern.Decorator_Pattern;

public class MushroomPizza extends ToppingDecorator{
    
    public MushroomPizza(BasePizza pizza){
        super(pizza);
    }

    @Override
    public int cost(){
        return this.pizza.cost() + 200;
    }
}
