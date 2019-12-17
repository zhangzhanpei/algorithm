/**
 * 爬楼梯, 给定一个 n 阶的楼梯, 一次可以爬 1 阶或 2 阶, 求有多少种不同的爬法能爬到楼梯顶部
 * 第 n 阶楼梯, 可以由 n-1 阶爬 1 阶上来, 也可以由 n-2 阶爬 2 阶上来
 * 使用 dp[i] 表示爬到第 i 阶的爬法, 那么 dp[i] = dp[i - 1] + dp[i - 2], 即爬到 n-1 的方法数加上爬到 n-2 的方法数
 */
class Solution {
    public static void main(String[] args) {
        System.out.println(climbStairs(8));
    }

    public static int climbStairs(int n) {
        if (n == 1) {
            return 1;
        }

        int[] dp = new int[n];
        dp[0] = 1; // 爬到第一阶只有一种方法
        dp[1] = 2; // 爬到第二阶有两种方法
        for (int i = 2; i < n; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }

        return dp[n - 1];
    }
}
