/**
 * 给定一个int数组, 求总的汉明距离
 * 总距离 = 每个元素相应位 1 的个数乘以 0 的个数的总和
 */
package main

import "fmt"

func main() {
    totalHammingDistance([]int{4, 14, 2})
}

func totalHammingDistance(nums []int) int {
    ret, n := 0, len(nums)
    for i := 0; i < 32; i++ {
        count := 0
        for _, val := range nums {
            if (val >> uint(i)) % 2 == 1 {
                count++ //末位为 1 时 count++
            }
        }
        ret += count * (n - count) //1 的个数乘以 0 的个数
    }
    return ret
}
