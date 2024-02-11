package Design_Patterns.Multithreading;

class Account {

    private int accountNumber;
    private double balance;
    
    public Account(int accountNumber,double balance) {
        this.accountNumber = accountNumber;
        this.balance = balance;
    }

    public void withdraw(double amount){

        System.out.println("Withdrawing money from the account "+accountNumber);
        for (int i = 0; i < 3; i++) {
            System.out.println("Processing... of amount :" + amount);
            try{
                Thread.sleep(500);
            }catch(Exception e){
                System.out.println("Exception occured "+e);
            }
        }

        synchronized (this) {
            if (balance >= amount) {
                // Simulate some processing time
                try {
                    Thread.sleep(100);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }

                // Perform the withdrawal
                balance -= amount;
                System.out.println(Thread.currentThread().getName() + " withdrew $" + amount +
                        ". Remaining balance: $" + balance);
            } else {
                System.out.println(Thread.currentThread().getName() + " cannot withdraw $" + amount +
                        ". Insufficient funds. Current balance: $" + balance);
            }
        }
    
        for (int i = 0; i < 3; i++) {
            System.out.println("Post processing... of amount :" + amount);
            try{
                Thread.sleep(500);
            }catch(Exception e){
                System.out.println("Exception occured "+e);
            }
        }
        
    }
    
    public double getBalance(){
        return this.balance;
    }
}

class WithdrawalThread extends Thread {
    
    private Account account;
    private double amount;

    public WithdrawalThread(Account account, double amount) {
        this.account = account;
        this.amount = amount;
    }

    @Override
    public void run() {
        account.withdraw(amount);
    }
}

public class block_synchronization {
    public static void main(String[] args) {
        
        Account acc = new Account(234928923, 1000);

        WithdrawalThread t1 = new WithdrawalThread(acc, 300);
        WithdrawalThread t2 = new WithdrawalThread(acc, 800);

        t1.start();
        t2.start();
        
        try {
            // Wait for both threads to complete before proceeding
            t1.join();
            t2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Current balance : " + acc.getBalance());
    }
}
