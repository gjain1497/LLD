package Design_Patterns.Structural_Pattern.Decorator_Pattern;

public abstract class ToppingDecorator extends BasePizza{
    protected BasePizza pizza;
    public ToppingDecorator(BasePizza p) { 
        this.pizza = p;
    }
}
