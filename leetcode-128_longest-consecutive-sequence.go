/**
 * 给定一个无序int数组, 返回最长的连续元素个数
 * 如 [100, 4, 200, 1, 3, 2] 的连续元素为[1, 2, 3, 4], 则返回长度4
 * 要求时间复杂度O(n)
 * 无序数组并且O(n), 优先考虑map
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
    m := map[int]bool{}
    for _, val := range nums {
        m[val] = true
    }
    longest := 0
    for _, val := range nums {
        length := 1
        for i := val - 1; m[i] == true; i-- {
            length++
        }
        for j := val + 1; m[j] == true; j++ {
            length++
        }
        longest = int(math.Max(float64(longest), float64(length)))
    }
    return longest
}
