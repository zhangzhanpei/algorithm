/**
 * 给定一个n元素的int数组, 所有元素1<=a[i]<=n, 其中有一些元素出现了两次, 找到这些元素
 * 遍历数组, 以元素值为下标将对应位置的元素改为负数.
 * 当遇到重复元素时, 相应位置的元素已经是负数了
 */
package main

import "fmt"
import "math"

func main() {
    findDuplicates([]int{4,3,2,7,8,2,3,1})
}

func findDuplicates(nums []int) []int {
    res := []int{}
    for _, val := range nums {
        idx := int(math.Abs(float64(val))) - 1 //因为元素最大为n, 所以改变n-1位置的元素, 否则会下标越界
        if nums[idx] < 0 { //如果当前元素作为下标, 对应位置的元素已是负数, 则当前元素重复了
            res = append(res, int(math.Abs(float64(val))))
        } else {
            nums[idx] = -nums[idx]
        }
        fmt.Println(nums)
    }
    return res
}
