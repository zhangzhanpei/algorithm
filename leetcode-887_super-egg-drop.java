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
            // 扔一次
            // dp[step - 1][i - 1] 鸡蛋碎了，剩下 i - 1个蛋和 step - 1 次机会去判断下面楼层数
            // dp[step - 1][i] 鸡蛋没碎，剩下 i 个蛋和 step - 1 次机会继续去判断上面楼层
            // 那么在 i 个蛋和 step 次机会下能判断的总楼层为 下面楼层 + 当前楼层 + 上面楼层
            dp[step][i] = dp[step - 1][i - 1] + dp[step - 1][i] + 1;
        }
    }
    return step;
}

// 空间复杂度 O(K*N) 时间复杂度O(K*N*N)
public int superEggDrop(int K, int N) {
    int[][] step = new int[K + 1][N + 1];
    for (int i = 1; i <= K; i++) {
        for (int j = 1; j <= N; j++) {
            step[i][j] = j;
        }
    }

    for (int i = 2; i <= K; i++) {
        for (int j = 2; j <= N; j++) {
            for (int k = 1; k < j; k++) { // 尝试次数应小于当前最大楼层数j
                // 在k层扔，有两种情况
                // 鸡蛋碎了，尝试次数为k层以下的次数加上当前这次 step[i - 1][k - 1] + 1
                // 鸡蛋没碎，尝试次数为k层往上至当前最大楼层j， step[i][j - k] + 1
                step[i][j] = Integer.min(step[i][j], Integer.max(step[i - 1][k - 1] + 1, step[i][j - k] + 1));
            }
        }
    }
    return step[K][N];
}
