/**
 * 给定一个 m*n 整数矩阵, 从左上角走到右下角(每步只能向右或向下走), 求经过的路径的元素的最小和
 * dp[i][j] = grid[i][j] + min(dp[i - 1][j], dp[i][j - 1])
 */

package main

import "fmt"
import "math"

func main() {
    minPathSum([][]int{
        {1, 2, 3},
        {4, 5, 6},
    })
}

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    var dp [][]int
    for i := 0; i < m; i++ {
        dp = append(dp, make([]int, n))
    }
    dp[0][0] = grid[0][0]
    for i := 1; i < m; i++ {
        dp[i][0] = grid[i][0] + dp[i - 1][0] //第一列一直向下加
    }
    for j := 1; j < n; j++ {
        dp[0][j] = grid[0][j] + dp[0][j - 1] //第一行一直向右加
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = grid[i][j] + int(math.Min(float64(dp[i - 1][j]), float64(dp[i][j - 1])))
        }
    }
    return dp[m - 1][n - 1]
}
