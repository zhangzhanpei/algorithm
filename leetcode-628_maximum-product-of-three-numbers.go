/**
 * 给定一个int数组, 取其中3个元素使得乘积最大, 返回最大乘积
 */
package main

import "fmt"
import "math"
import "sort"

func main() {
    fmt.Println(maximumProduct([]int{1,2,3,4}))
}

func maximumProduct(nums []int) int {
    sort.Ints(nums) //排序
    a := nums[len(nums) - 1] * nums[len(nums) - 2] * nums[len(nums) - 3] //最大的3个元素相乘
    b := nums[0] * nums[1] * nums[len(nums) - 1] //最左边如果是两个负数, 相乘后乘以最大元素
    return int(math.Max(float64(a), float64(b))) //取两者中的大者
}
