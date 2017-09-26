/**
 * 给定一个无序int数组, 求升序子序列最大长度
 * nums = [10, 9, 2, 5, 3, 7, 101, 18], 则最长升序子序列为 [2, 3, 7, 101], return 4
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
    n := len(nums)
    ret := 0
    dp := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = 1
    }
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = int(math.Max(float64(dp[i]), float64(dp[j] + 1)))
            }
        }
        ret = int(math.Max(float64(ret), float64(dp[i])))
    }
    return ret
}
