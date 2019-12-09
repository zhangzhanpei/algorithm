import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * 大数阶乘。给定一个整数n, 求n!
 * 注意 n > 12 时 int 的溢出问题
 */
class Solution {
    public static void main(String[] args) {
        factorial(10);
    }

    public static void factorial(int n) {
        // 使用列表倒序存储阶乘结果
        List<Integer> ret = new ArrayList<>();
        ret.add(1);
        for (int i = 2; i <= n; i++) {
            // 进位
            int carry = 0;
            // 列表每个元素都乘以下一个阶乘数
            for (int j = 0; j < ret.size(); j++) {
                // 乘积加上进位
                int tmp = ret.get(j) * i + carry;
                carry = tmp / 10;
                ret.set(j, tmp % 10);
            }
            // 处理进位。这里的进位是乘出来的，例如 9 * 100，不是相加那种只进一位的，所以要循环处理
            while (carry > 0) {
                ret.add(carry % 10);
                carry /= 10;
            }
        }

        // 反转列表
        Collections.reverse(ret);
        // 遍历输出即 n 的阶乘结果
        for (Integer i : ret) {
            System.out.print(i);
        }
    }
}
