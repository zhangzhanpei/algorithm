//给定一个int数组, 其中一个子集的和最大, 求这个和
package main

import "fmt"

func main() {
    fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
    max := nums[0]
    var sum int
    for i := 0; i < len(nums); i++ {
        //如果sum加上当前元素还不如当前元素大, 那么子集从当前元素重新开始计算, 因为加上前面的元素反而sum更小
        if sum + nums[i] <= nums[i] {
            sum = nums[i]
        } else {
            sum += nums[i]
        }

        if sum > max {
            max = sum
        }
    }
    return max
}
