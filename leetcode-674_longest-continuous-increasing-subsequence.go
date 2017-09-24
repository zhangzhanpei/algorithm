/**
 * 给定一个int数组, 求连续增长的最大子序列长度
 */
/**
 * 给定一个int数组, 求最长升序子序列长度
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(findLengthOfLCIS([]int{2,2,2}))
}

func findLengthOfLCIS(nums []int) int {
    max, tmp := 0, 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] > nums[i - 1] {
            tmp++
            max = int(math.Max(float64(max), float64(tmp)))
        } else {
            tmp = 1
        }
    }
    return max
}
