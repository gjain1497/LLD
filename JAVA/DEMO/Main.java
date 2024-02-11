package DEMO;

public class Main {
    
    public static void main(String[] args) throws Exception{

        int points = 10000000;
        int pointInsideCircle = 0;

        for(int i=0;i<points;i++){

            double x = Math.random();
            double y = Math.random();
            
            double val = Math.pow(x-0.5,2) + Math.pow(y-0.5, 2);
            double distance = Math.sqrt(val);

            if(distance <= 0.5){
                pointInsideCircle++;
            }
        }
        double PI = (double)4*pointInsideCircle / (double)points;
        System.out.println("The approximate value of PI : "+PI);




        // int points = 1000000;
        // double PI = 3.141;
        // int lp = 0, hp = 1000000;
        // double currPI = 0;
        // while(Math.abs(PI-currPI)  > 0.000001){

        //     int pointInsideCircle = (lp+hp)/2;
        //     currPI = (double)4*pointInsideCircle / (double)points;

        //     if(currPI > PI) {
        //         hp = pointInsideCircle-1;
        //     }
        //     else {
        //         lp = pointInsideCircle+1;
        //     }
        // }
        // System.out.println("The approximate value of PI : "+currPI);

    }
}


