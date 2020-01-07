/**
 * 给定一个 int 数组，判断能否拆成两部分使得元素的和相等
 * 数组的和必须为偶数，才能使拆出来的两部分和相等
 * 问题转换为从数组中取元素使得和等于原数组的一半
 * 0-1 背包问题
 */
public boolean canPartition(int[] nums) {
    int sum = 0;
    for (int i : nums) {
        sum += i;
    }
    // 如果数组的和不是偶数，说明无法拆分
    if (sum % 2 != 0) {
        return false;
    }
    // 目标值为和的一半
    sum /= 2;
    // 定义 dp[i] 为前 i 个元素能否得到和为 i
    boolean[] dp = new boolean[sum + 1];
    dp[0] = true;
    for (int num : nums) {
        for (int j = sum; j >= num; j--) {
            dp[j] = dp[j] || dp[j - num];
        }
    }
    return dp[sum];
}
