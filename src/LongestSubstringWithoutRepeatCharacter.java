/* Created by Fitz on 2017/3/19 */

import java.util.HashMap;

public class LongestSubstringWithoutRepeatCharacter {
    public static void main(String[] args) {
        LongestSubstringWithoutRepeatCharacter lswrc = new LongestSubstringWithoutRepeatCharacter();
        int a = lswrc.solution("asdaswdsas");
        System.out.println(a);
    }

    public int solution(String s){
        if (s.length() == 0)
            return 0;
        int max = 0;
        HashMap<Character,Integer> map = new HashMap<Character, Integer>();
        for (int i = 0, j = 0; i < s.length(); i++) {
            if (map.containsKey(s.charAt(i))){
               j = Math.max(j,map.get(s.charAt(i)) + 1);
            }
            map.put(s.charAt(i),i);
            max = Math.max(max,i-j+1);
        }
        return max;
    }
}
