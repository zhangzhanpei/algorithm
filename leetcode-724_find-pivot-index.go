/**
 * 给定一个int数组，找到数组中心点[左右两侧的元素之和相等]的下标
 */
package main

import "fmt"

func main() {
    fmt.Println(pivotIndex([]int{-1,-1,-1,0,1,1}))
}

func pivotIndex(nums []int) int {
    sum, left := 0, 0
    for _, v := range nums { //求和
        sum += v
    }
    for i := 0; i < len(nums); i++ {
        if i > 0 { //下标0也可以是中间点，因为右侧元素之和可能为0
            left += nums[i - 1]
        }
        if left * 2 + nums[i] == sum { //如果下标i两侧的和相等，即i为中点位置
            return i
        }
    }
    return -1
}
