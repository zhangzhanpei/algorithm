import java.util.Stack;

/**
 * 给定一个字符串包含相同数量的R和L，切割该字符串使得每个字串的R和L数量仍一致，求最多能切成多少个字串
 */
class Solution {
    public static void main(String[] args) {
        String s = "RLRRLLRLRL";
        System.out.println(balancedStringSplit(s));
    }

    public static int balancedStringSplit(String s) {
        int ret = 0;
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (!stack.empty() && stack.peek() != c) {
                stack.pop();
                if (stack.empty()) {
                    ret++;
                }
            } else {
                stack.push(c);
            }
        }
        return ret;
    }
}
