/**
 * 小偷沿街偷东西, 不能偷相邻的两家, 求能偷到的最大价值
 * 即给定一个数组, 取出任意不相邻的元素使得和最大
 * 动态规划
 * dp[i] 表示 0~i 能得到的最大和, 那么
 * dp[0] = nums[0] //只有一家能偷
 * dp[1] = max(nums[0], nums[1]) //0和1相邻, 只能偷一个
 * dp[i] = max(dp[i-1], dp[i-2] + nums[i]) //dp[i-1]即不偷第i家, dp[i-2]+dp[i]即偷第i家, 两者取大的那一个
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(rob([]int{4, 3, 1, 3, 2}))
}

func rob(nums []int) int {
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
