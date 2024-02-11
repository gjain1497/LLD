package LLD.TicTacToe;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

public class TTTGame {
    
    private int emptySlots;
    private int boardSize; 
    private List<Player>players = new ArrayList<Player>();
    private  boolean currentPlayer; //the player whose turn it is currently
    private char[][] board; //2D array to represent the game board
    private HashSet<String>winConditions = new HashSet<String>();

    private GameState status = GameState.NOT_OVER;
    private Player winner; 

    public TTTGame(int boardSize, Player p1, Player p2){
        players.add(p1);
        players.add(p2);
        this.boardSize = boardSize;
        board = new char[boardSize][boardSize];
        configureBoard();
        currentPlayer = false; //start with player 1's turn
    }

    private  void configureBoard(){

        emptySlots = boardSize*boardSize;

        // set up initial configuration
        for(int i=0;i<boardSize;i++){
            for(int j=0;j<boardSize;j++){
                board[i][j] = (char)('0'+i*boardSize+j);
            }
        }
        
        // set up win conditions for a standard tic-tac-toe game (n*n)
        for(Player p : players){
            char mark = p.getMark();
            String temp = "";
            for(int i=0;i<boardSize;i++){
                temp += mark;
            }
            winConditions.add(temp);
        }
    }

    public void showBoard(){

        System.out.println("Current Game Board:");
        for(int i=0;i<boardSize;i++){
            for(int j=0;j<boardSize;j++){
                System.out.print(board[i][j]+" ");
            }
            System.out.println();
        }
        System.out.println();
    }

    public Player getCurrPlayer(){
        return players.get(currentPlayer ? 1 : 0);
    }
    public void setCurrentPlayer(){
        currentPlayer = !currentPlayer;
    }
    public boolean isGameOver(){
        return status == GameState.OVER || status == GameState.TIE;
    }
    public boolean isTie(){
        return status == GameState.TIE;
    }
    public void setGameStatus(GameState status){
        this.status = status;
    }
    public Player getWinner() {
        return winner;
    }
    public void setWinner(Player p){
        this.winner = p;
    }

    public void makeMove(int slot) throws Exception{

        int row = slot/boardSize ,col = slot%boardSize;

        if(!validSlot(row, col)){
            throw new Exception("Out of boundary");
        }else if(board[row][col] == 'O' || board[row][col] == 'X'){
            throw new Exception("Slots already occupied");
        }

        board[row][col] = getCurrPlayer().getMark();

        if (checkForWin(row,col)) {
            this.winner = getCurrPlayer();
            setGameStatus(GameState.OVER);
        }else if(isBoardFull()){
            setGameStatus(GameState.TIE);
        }

        setCurrentPlayer();
        emptySlots--;

    }

    private boolean validSlot(int row,int col){
        if(row < 0 || row >= boardSize || col < 0 || col >= boardSize )
            return false;
        return true;
    }
    private boolean isBoardFull(){
        return emptySlots <= 0;
    }

    private boolean checkForWin(int row,int col){
        
        String temp;

        // Check rows
        if(checkRowWin(row)) 
            return true;

        // Check for column
        if(checkColWin(col)) 
            return true;

        // Check for diagonal 
        if(row == col)
        {
            temp = "";
            for(int i=0;i<boardSize;i++) 
                temp += board[i][i];
            if(isWinCondition(temp))  return true;

            temp = "";
            for(int i=boardSize-1;i>=0;i--) 
                temp += board[i][i];
            if(isWinCondition(temp))  return true;
        }

        return false;
            
    }
    private boolean checkRowWin(int row){

        String temp = "";
        for(int j=0;j<boardSize;j++)
            temp += board[row][j];
        if(temp.length() > 0 && isWinCondition(temp)){
            return true;
        }
        return false;
    }

    private boolean checkColWin(int col){
        String temp = "";
        for(int i=0;i<boardSize;i++)
            temp += board[i][col];
        if(temp.length() > 0 && isWinCondition(temp)){
            return true;
        }
        return false;
    }
    private boolean isWinCondition(String temp) {

        if(winConditions.contains(temp))
          return true;
        
        return false;
    }


}
