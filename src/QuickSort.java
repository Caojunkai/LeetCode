/* Created by Fitz on 2017/3/23 */

public class QuickSort {

    public void solution(int[] arr,int left, int right){
        if (left >= right){
            return;
        }
        int i = left;
        int j = right;
        int key = arr[i];
        while (i < j){
            while (i < j && key <= arr[j]){
                j--;
            }
            arr[i] = arr[j];
            while (i < j && key >= arr[i]){
                i++;
            }
            arr[j] = arr[i];
            arr[i] = key;
        }
        solution(arr,left,i-1);
        solution(arr,i+1,right);
    }
}
