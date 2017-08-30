/**
 * 给定一个m*n的矩阵, 从左上角开始走, 只能向右或向下走, 求走到最右下角有几种不同的走法
 * 动态规划
 * path[i][j] 表示走到ij位置的不同走法, 则有
 * path[i][j] = path[i][j - 1] + path[i - 1][j] 走到ij的走法等于走到上一格的方法加上走到左边一格的走法数
 */
func uniquePaths(m int, n int) int {
    var dp [][]int
    for i := 0; i < m; i++ {
        arr := make([]int, n)
        for j := 0; j < n; j++ {
            arr[j] = 1
        }
        dp = append(dp, arr)
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        }
    }
    return dp[m - 1][n - 1];
}
