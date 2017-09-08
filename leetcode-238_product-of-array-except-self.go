/**
 * 给定一个int数组, 遍历数组, 返回除当前元素外其他元素的乘积数组(不使用除法)
 */
package main

import "fmt"

func main() {
    productExceptSelf([]int{1, 2, 3, 4})
}

func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := []int{1}
    for i := 1; i < n; i++ {
        res = append(res, res[i - 1] * nums[i - 1])
    }
    right := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= right;
        right *= nums[i];
    }
    return res
}
