//给定n个元素的int数组, 从其中连续取k个元素使得平均值最大
package main

import "fmt"

func main() {
    fmt.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ { //求数组前k个元素的和
        sum += nums[i]
    }
    max := sum
    for i := k; i < len(nums); i++ { //移动窗口, sum加上下一个元素并减去第一个元素, 此时与之前的sum比较
        sum += nums[i] - nums[i - k]
        if sum > max {
            max = sum
        }
    }
    return float64(max) / float64(k)
}
