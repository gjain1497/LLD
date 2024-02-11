package Design_Patterns.Multithreading;

import java.util.HashMap;

class CacheManager {

    private static HashMap<String, String> cache = new HashMap<>();

    public static String getCache(String key) {
        return cache.get(key);
    }

    public static void setCache(String key, String value) {
        cache.put(key, value);
    }
}

class CacheThread extends Thread {
    @Override
    public void run() {
        
        String result;
        synchronized(CacheManager.class){
            result = CacheManager.getCache("data");
            if (result == null) {
                System.out.println("Data is not in the cache, " + Thread.currentThread().getName() + " fetching from database -----");
                result = fetchFromDatabase();
                CacheManager.setCache("data", result);
            }
        }
        System.out.println(Thread.currentThread().getName() + " got value: " + result);
    }

    public String fetchFromDatabase() {
        System.out.println("Reading from database -----");
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            System.out.println(e);
        }
        return "Some data from the database";
    }
}

public class class_synchronization {
    public static void main(String[] args) {

        CacheThread t1 = new CacheThread();
        CacheThread t2 = new CacheThread();

        t1.start();
        t2.start();

        try {
            t1.join();
            t2.join();
        } catch (Exception e) {
            System.out.println(e);
        }
        System.out.println("Main terminates");
    }
}
