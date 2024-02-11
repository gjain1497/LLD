package Design_Patterns.demo;

import java.util.Arrays;
import java.util.Scanner;

public class Main {
    public static void main(String[] args){

        System.out.println("Hello");

        String name = "Kartik";
        int age = 22;
        float pi = 3.14F;
        System.out.println(name + " : " +  age);
        System.out.println(pi);

        String name2 = name.substring(2,4);
        System.out.print(name2);


        System.out.println("1D arrays");
        int[] marks  = new int[] {4,3,2};
        for(int num : marks){
            System.out.print(num + " ");
        }
        System.out.println();

        System.out.println(marks.length);
        // Sort 
        Arrays.sort(marks);
        System.out.println("After sorting");
        for(int i=0;i<marks.length;i++){
            System.out.print(marks[i] + " ");
        }
        System.out.println();

        System.out.println("2D arrays");
        int[][] finalMarks = {{45,67,89},{11,234,45}};
        for (int[] finalMark : finalMarks) {
            for (int mark : finalMark) {
                System.out.print(mark + " ");
            }
            System.out.println();
        }
        System.out.println();

        //constants
        final  double PI = 3.141592653;
        System.out.println(PI);

        //Maths class
        System.out.println(Math.max(2,3));
        System.out.println((int)(Math.random()*10));

        //Inputs in Java
        Scanner sc = new Scanner(System.in);


        //Take input as int 
        // int val = sc.nextInt();
        // float x = sc.nextFloat();
        // System.out.println(val + " "+x);

        //Take input as string
        // String fullName = sc.nextLine();
        // System.out.println("Your name is "+ fullName);

        sc.close();
        int y = 10;

        switch (y) {
            case 2:
                System.out.println("Value is 2");
                break;
            case 10:
                System.out.println("Value is 10");
                break;
            default:
                System.out.println("No cases  matched!");
                break;
        }

        // Exception Handling
        try {
            int a = 7/1;
            System.out.println(a);
        } catch (Exception e) {
            System.out.println("An error occured! ");
        } 
        return;
    }
}
