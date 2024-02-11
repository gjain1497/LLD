package Design_Patterns.demo;

class Pen {
    String color;
    String type;

    public Pen(){

    }
    public Pen(String color , String type){
        this.color = color;
        this.type = type;
    }
    public void  draw() {
        System.out.println("Write something");
    }
}

class Student {
    String name;
    int age,roll;
    static String school = "DPS";
    Student(){

    }
    Student(String name,int age,int roll){
        this.name = name;
        this.age = age;
        this.roll = roll;
    }
    public void displayInfo(){
        System.out.println(this.name + " " + Student.school + " " + this.age + " " + this.roll);
    }
    public static void changeSchool(String school){
        Student.school = school;
    }
    
}

public class Oops {
    public static void main(String[] args) {
        
        Student s1 = new Student("James",12,345635657);
        Student s2  = new Student("Tony",35,90245728);

        s1.displayInfo();
        s2.displayInfo();

        System.out.println("After changing school");
        Student.changeSchool("DAV");
        s1.displayInfo();
        s2.displayInfo();
    }
}
