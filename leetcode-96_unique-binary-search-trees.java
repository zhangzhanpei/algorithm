/**
 * 给定整数n, 求能构建多少个不同的二叉搜索树(节点大小为1 <= i <= n)
 * n = 0 时为空树, return 1
 * n = 1 时, return 1
 * n = 2 时, return 2
 * n > 2 时, dp[n]= dp[0] * dp[n-1] + dp[1] * dp[n-2] + ... + dp[n-1] * dp[0]
 */
public int numTrees(int n) {
    int[] dp = new int[n + 1];
    dp[0] = dp[1] = 1;
    for (int i = 2; i <= n; i++) {
        for (int j = 0; j < i; j++) {
            dp[i] += dp[j] * dp[i - j - 1];
        }
    }
    return dp[n];
}
