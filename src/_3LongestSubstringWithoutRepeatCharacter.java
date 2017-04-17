/* Created by Fitz on 2017/3/19 */

import java.util.HashMap;
import java.util.Map;

public class _3LongestSubstringWithoutRepeatCharacter {
    public static void main(String[] args) {
        _3LongestSubstringWithoutRepeatCharacter lswrc = new _3LongestSubstringWithoutRepeatCharacter();
        String s = "asdaswdsassdasdasdasdqwqeqweqwe";
        int a = lswrc.solution(s);
        System.out.println(a);
    }


    public int solution(String s){
        if (s == null || s.length() == 0)
            return 0;
        int leftBound = 0,max = 0;
        Map<Character,Integer> map = new HashMap<Character, Integer>();
        for (int i = 0; i < s.length(); i++) {
            if (map.containsKey(s.charAt(i))){
                leftBound = Math.max(leftBound,map.get(s.charAt(i))+1);
            }
            map.put(s.charAt(i),i);
            max = Math.max(max,i-leftBound+1);
        }
        return max;
    }
}
