/**
 * 给定一个m*n的矩阵, 从左上角开始走, 只能向右或向下走, 求走到最右下角有几种不同的走法(矩阵中为1的位置即障碍物不能走)
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
        return 0
    }
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([]int, n)
    dp[0] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
                continue
            }
            if j > 0 {
                dp[j] += dp[j - 1]
            }
        }
    }
    return dp[n - 1]
}
