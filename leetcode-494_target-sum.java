/**
 * 给定一个int数组和一个目标值，数组元素之间只能使用+和-号使得和等于目标值，求有多少种解法
 */
public int findTargetSumWays(int[] nums, int S) {
    int sum = 0;
    for (int n : nums) {
        sum += n;
    }
    if (sum < S || (S + sum) % 2 != 0) {
        return 0;
    }
    return this.subsetSum(nums, (S + sum) / 2);
}

public int subsetSum(int[] nums, int s) {
    int[] dp = new int[s + 1];
    dp[0] = 1;
    for (int n : nums) {
        for (int i = s; i >= n; i--) {
            dp[i] += dp[i - n];
        }
    }
    return dp[s];
}
