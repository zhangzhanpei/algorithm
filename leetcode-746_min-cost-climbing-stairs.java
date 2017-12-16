/**
 * 给定一个int数组表示跳每级台阶的花费，一次可跳一阶或两阶，求跳到台阶顶部的最小花费
 */
public static int minCostClimbingStairs(int[] cost) {
    int len = cost.length;
    int[] dp = new int[len + 1];
    dp[0] = cost[0];
    dp[1] = cost[1];

    for (int i = 2; i <= len; i++) {
        int currentCost = i == len ? 0 : cost[i];
        dp[i] = Math.min(dp[i - 1], dp[i - 2]) + currentCost; //跳最后一阶的花费是固定的，因此只需选择前一阶或前两阶花费小的那个跳法
    }
    return dp[len];
}
