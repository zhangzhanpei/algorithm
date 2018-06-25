/**
 * 完全背包问题
 * 一个固定容量的背包，一堆数量无限的物品，如何使装进背包的物品价值最大
 * 这个和 0-1 背包的区别就在于每种物品的数量是无限制的
 */
class Solution {
    public static void main(String[] args) {
        // 每个物品的体积
        int[] volume = {2, 3, 4, 7};
        // 每个物品的价值
        int[] value = {1, 3, 5, 9};
        // 背包总容量
        int capacity = 10;

        System.out.println(knapsack(volume, value, capacity));
    }

    public static int knapsack(int[] volume, int[] value, int capacity) {
        int[] dp = new int[capacity + 1];
        for (int i = 0; i < volume.length; i++) {
            for (int j = 0; j <= capacity; j++) {
                // 对每个物品有不取，取一件和多件等方案
                // 当背包剩余容量大于当前物品时，继续装当前物品
                // dp[j] 指背包容量为 j 时，前 i 种物品能取得的最大价值
                if (j >= volume[i]) {
                    dp[j] = Integer.max(dp[j], dp[j - volume[i]] + value[i]);
                }
            }
        }
        return dp[capacity];
    }
}
