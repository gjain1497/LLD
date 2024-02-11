package Design_Patterns.Structural_Pattern.Proxy_Pattern.Subject;

public class Proxy extends RealSubject{
    
    public void method(){
        
        System.out.println("This is the proxy's implementation of the method()");

        System.out.println(" If authenticated successfully");
        super.method();
    }
}
