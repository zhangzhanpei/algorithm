/**
 * 给定 K 个鸡蛋和 N 层高的大楼，求最少的次数判断最高哪一层扔鸡蛋不会碎
 */
public int superEggDrop(int K, int N) {
    // 定义 dp[step][K] 为 K 个鸡蛋在 step 次数时能检查的最大楼层数 F
    int[][] dp = new int[N + 1][K + 1];
    int step = 0;
    while (dp[step][K] < N) {
        step++;
        for (int i = 1; i <= K; i++) {
            dp[step][i] = dp[step - 1][i - 1] + dp[step - 1][i] + 1;
        }
    }
    return step;
}
