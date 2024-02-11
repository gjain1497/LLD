package Design_Patterns.Structural_Pattern.Proxy_Pattern;

import Design_Patterns.Structural_Pattern.Proxy_Pattern.Subject.*;


public class Main {
    public static void main(String[] args) {
        Subject obj = new Proxy();
        obj.method();
    }
}
