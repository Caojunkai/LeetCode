/* Created by Fitz on 2017/3/15 */

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;

public class _1TwoSum {
    public static void main(String[] args) {
        int[] numbers = new int[]{2,7,11,15};
        int target = 22;
        _1TwoSum twoSum1 = new _1TwoSum();
        int[] result = twoSum1.solution(numbers,target);
        System.out.println(result[0]);
        System.out.println(result[1]);
        Deque<Integer> deque = new ArrayDeque<Integer>();
    }

    public int[] solution(int[] numbers, int target){
        int[] result = new int[2];
        Map<Integer,Integer> map = new HashMap<Integer, Integer>();
        for (int i = 0; i < numbers.length; i++){
            if (map.containsKey(target - numbers[i])){
                result[1] = i;
                result[0] = map.get(target - numbers[i]);
                return result;
            }
            map.put(numbers[i],i);
        }
        return result;
    }
}
