/**
 * 给定一个int数组，判断每个数字出现的次数是否都是不同的
 */

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;

class Solution {
    public static void main(String[] args) {
        int[] arr = {1, 2, 2};
        System.out.println(uniqueOccurrences(arr));
    }

    public static boolean uniqueOccurrences(int[] arr) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i : arr) {
            int count = map.containsKey(i) ? map.get(i) : 0;
            map.put(i, count + 1);
        }
        HashSet<Integer> set = new HashSet<>();
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            if (!set.add(entry.getValue())) {
                return false;
            }
        }
        return true;
    }
}
