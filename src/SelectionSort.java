/* Created by Fitz on 2017/3/23 */

public class SelectionSort {

    public void solution(int[] arr){
        int length = arr.length;
        for (int i = 0; i < length; i++){
            int min_index = 0;
            for (int j = i +1; j < length; j++){
                if (arr[j] < arr[min_index]){
                    min_index = j;
                }
            }
            int temp;
            temp = arr[i];
            arr[i] = arr[min_index];
            arr[min_index] = temp;
        }
    }
}
