/* Created by Fitz on 2017/3/22 */

public class BubbleSort {
    public static void main(String[] args) {
        int[] arr = {2,3,4,53,3432,2334,334,2423,4234,324,23};
        sort(arr);
        for (int item : arr){
            System.out.print(item + "->");
        }
    }

    public static void sort(int[] arr){
        int length = arr.length;
        for (int i = 0; i < length -1; i++) {
            for (int j = 0; j < length - i - 1; j++) {
                if (arr[j] > arr[j+1]){
                    int temp = arr[j];
                    arr[j] = arr[j+1];
                    arr[j+1] = temp;
                }
            }
        }
    }
}
