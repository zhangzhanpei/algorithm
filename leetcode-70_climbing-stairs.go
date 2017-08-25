/**
 * 爬楼梯, 给定一个n阶的楼梯, 一次可以爬1阶或2阶, 求有多少种不同的爬法能爬到楼梯顶部
 * 第n阶楼梯, 可以由n-1阶爬1阶上来, 也可以由n-2阶爬2阶上来
 * 使用dp[i]表示爬到第i阶的爬法, 那么dp[i] = dp[i - 1] + dp[i - 2], 即爬到n-1的方法数加上爬到n-2的方法数
 */
package main

import "fmt"

func main() {
    fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    dp := make([]int, n)
    dp[0] = 1 //爬到第一阶只有一种爬法
    dp[1] = 2 //爬到第二阶有两种爬法
    for i := 2; i < n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[n - 1]
}
