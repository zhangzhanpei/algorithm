import java.util.ArrayList;

/**
 * 多重背包问题
 * 一个固定容量的背包，一堆数量有限的物品，如何使装进背包的物品价值最大
 * 这个和 0-1 背包、完全背包的区别就在于每种物品的数量是有限制的
 */
class Solution {
    public static void main(String[] args) {
        // 每个物品的体积
        int[] volume = {1, 2, 2};
        // 每个物品的价值
        int[] value = {6, 10, 20};
        // 每个物品的数量
        int[] count = {10, 5, 2};
        // 背包总容量
        int capacity = 8;

        // 0-1 背包中物品是可以重复的，因此多重背包直接展开成重复物品，以 0-1 背包方案求解
        ArrayList<Integer> newVolume = new ArrayList<>();
        ArrayList<Integer> newValue = new ArrayList<>();
        for (int i = 0; i < volume.length; i++) {
            for (int j = 0; j < count[i]; j++) {
                newVolume.add(volume[i]);
                newValue.add(value[i]);
            }
        }

        System.out.println(knapsack(newVolume.stream().mapToInt(i -> i).toArray(), newValue.stream().mapToInt(i -> i).toArray(), capacity));
    }

    /**
     * 0-1 背包
     */
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
