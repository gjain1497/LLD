package Data_Strcutures;


public class Main {
    public static void main(String[] args) {

        // Arrays
        int[] arr = new int[] {2,3,4,5,6};
        for(int num : arr){
            System.out.println(num);
        }

        // 2D arrays
        int[][] nums = new int[][] {{1,2,3},{4,5,6}};
        for(int[] num : nums){
            for(int val : num){
                System.out.print(val + " ");
            }
            System.out.println();
        }

        // Jagged arrays
        int[][]  jaggedArr = new int[2][];
        jaggedArr[0] = new int[]{1,2,3};
        jaggedArr[1] = new int[] {4,5,6,7};

        for(int[] jag : jaggedArr){
            for(int num : jag){
                System.out.print(num + " ");
            }
            System.out.println();
        }


    }
}
