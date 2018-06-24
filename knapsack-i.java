/**
 * 0-1 背包问题
 * 一个固定容量的背包，一堆物品，如何使装进背包的物品价值最大
 */
class Solution {
    public static void main(String[] args) {
        // 每个物品的体积
        int[] volume = {3, 5, 2, 6, 4};
        // 每个物品的价值
        int[] value = {4, 4, 3, 5, 3};
        // 背包总容量
        int capacity = 12;

        System.out.println(knapsack(volume, value, capacity));
    }

    public static int knapsack(int[] volume, int[] value, int capacity) {
        int[] dp = new int[capacity + 1];
        for (int i = 0; i < volume.length; i++) {
            for (int j = capacity; j >= volume[i]; j--) {
                // 对每个物品有取和不取两种方案
                // 不取物品，容量和价值都不变
                // 取物品时，容量减少，价值增加
                dp[j] = Integer.max(dp[j], dp[j - volume[i]] + value[i]);
            }
        }
        return dp[capacity];
    }
}
