package LLD.TicTacToe;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        System.out.println("=== Set up some configuration before game begins ===");
        Scanner sc = new Scanner(System.in);

        System.out.println("Player 1, please enter your name:");
        String name = sc.nextLine();
        Player p1 = new Player((int)(Math.random()*1000), name);

        System.out.println("Player 2, please enter your name:");
        name = sc.nextLine();
        Player p2 = new Player((int)(Math.random()*1000), name);

        p1.setMark('X');
        p2.setMark('O');
        
        System.out.println("Enter size of the board");
        int boardSize = Integer.parseInt(sc.nextLine()); // 3,4 or
        TTTGame game = new TTTGame(boardSize,p1,p2);
       

        System.out.println("=== Game Starts ===");
        System.out.println("-- Choose a numbered slot to play --");

        while (!game.isGameOver()) {
            game.showBoard();
            System.out.println(game.getCurrPlayer().getName() + "'s turn.");
            while(true){
                try {
                    int move = sc.nextInt();
                    game.makeMove(move);
                    break;
                } catch (Exception e) {
                    System.out.println(e.getMessage());
                    System.out.println("Invalid move! Please choose again.");
                }
            }
        }
        System.out.println("\n\n === Game Over === ");
        game.showBoard();
        if(game.isTie())
            System.out.println("It's a tie!");
        else
            System.out.println(game.getWinner().getName()+" wins!");
        
        sc.close();
    }
}
