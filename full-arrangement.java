import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * 给定一个字符数组, 返回字符的所有排列方式
 */
class Solution {
    private char[] chars = {'a', 'b', 'c', 'd'}; // 可使用的排列字符
    private char[] str = new char[4]; // 存储一种排列
    private Map<Character, Boolean> used = new HashMap<>(); // 标记某个字符是否已使用

    public void fullArrangement(int step) {
        if (step == 4) {
            System.out.println(Arrays.toString(str)); // 输出一种排列
            return;
        }

        for (int i = 0; i < 4; i++) {
            if (!used.getOrDefault(chars[i], false)) {
                str[step] = chars[i];
                used.put(chars[i], true);
                fullArrangement(step + 1);
                // 当输出一种排列后返回, 取回当前 step 尝试的字符, 换成下一个字符试试.
                // 即每一个 step 都会尝试所有的字符
                used.put(chars[i], false);
            }
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        solution.fullArrangement(0); // 从第一个位置开始尝试
    }
}
