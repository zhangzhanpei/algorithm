/**
 * 给定一个 int 数组表示硬币面额，以及一个总金额 amount，求凑成总金额的最少硬币数
 */
public int coinChange(int[] coins, int amount) {
    // dp[i] 代表凑成 i 元时的最小硬币数
    int[] dp = new int[amount + 1];
    for (int i = 1; i <= amount; i++) {
        dp[i] = Integer.MAX_VALUE;
        for (int coin : coins) {
            // 金额大于当前硬币金额才进入取不取当前硬币的判断
            if (i >= coin) {
                // 取当前硬币，则待凑金额变小，硬币数+1
                dp[i] = Math.min(dp[i], dp[i - coin] + 1);
            }
        }
    }
    return dp[amount] == Integer.MAX_VALUE ? -1 : dp[amount];
}
