/* 
 * 给定 n 个元素的 int 数组, 每次 n-1 个元素 +1, 需要加多少次才能让所有元素相等
 * 假设需要 m 次, 最后所有元素的值为 x, 一开始数组的最小值为 minNum, 和为 sum, 即
 * sum + m * (n - 1) = x * n
 * 另外有 minNum 经过 m 次 +1 后得到了 x
 * x = minNum + m
 * 可推导出
 * sum - minNum * n = m
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(minMoves([]int{1, 2, 3}))
}

//sum(arr) - len(arr) * min(arr)
func minMoves(nums []int) int {
    sum := 0
    minNum := math.MaxInt64
    for _, v := range nums {
        sum += v //求和
        if v < minNum {
            minNum = v //求最小值
        }
    }
    return sum - len(nums) * minNum
}
