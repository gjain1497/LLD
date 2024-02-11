package LLD.TicTacToe;

public class Player {
    
    private final int ID;
    private String name;
    private char mark;

    Player(int ID,String name){
        this.ID = ID;
        this.name = name;
    }
    public String getName(){
        return name;
    }
    char getMark(){
        return this.mark;
    }
    void setMark(char mark) {
        this.mark = mark;
    }

}
