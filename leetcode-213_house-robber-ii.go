/**
 * 参照leetcode-198
 * 小偷沿街偷东西, 不能偷相邻的两家, 求能偷到的最大价值(注意街道是环形的, 偷了第一家就不能偷最后一家)
 * 偷第一家和不偷第一家的价值比较取大者即可
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(rob([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}))
}

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    //偷第一家, 不偷最后一家
    //不偷第一家, 偷最后一家
    //两者比较
    return int(math.Max(float64(robber(nums[:len(nums) - 1])), float64(robber(nums[1:len(nums)]))))
}

func robber(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    size := len(nums)
    dp := make([]int, size)
    dp[0] = nums[0]
    dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
    for i := 2; i < size; i++ {
        dp[i] = int(math.Max(float64(dp[i - 1]), float64(dp[i - 2] + nums[i])))
    }
    return dp[size - 1]
}
