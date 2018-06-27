import java.util.Stack;

/**
 * 给定两个字符串 S 和 T，字符串中的 # 代表退格键，判断两个字符串最终是否一致
 */
class Solution {
    public static boolean backspaceCompare(String S, String T) {
        return backspace(S).equals(backspace(T));
    }

    public static String backspace(String s) {
        // 使用栈模拟字符串的输入过程
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '#') {
                if (!stack.empty()) {
                    stack.pop();
                }
            } else {
                stack.push(s.charAt(i));
            }
        }
        // 得到最终字符串
        StringBuilder ret = new StringBuilder();
        for (char c : stack) {
            ret.append(c);
        }
        return ret.toString();
    }
}
