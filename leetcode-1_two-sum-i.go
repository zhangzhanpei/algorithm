/**
 * 给定一个有序int数组和一个target, 从数组中取两个元素使得和等于target
 */
package main

import "fmt"

func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
}

func twoSum(nums []int, target int) []int {
    i, j := 0, len(nums) - 1
    for i < j {
        //因为是有序数组, 从数组两头开始找
        sum := nums[i] + nums[j]
        if sum > target { //sum > target 时需要减小 sum, 即右边向前取更小的元素
            j--
        } else if sum < target { //sum < target 时需要加大 sum, 即左边向后取更大的元素
            i++
        } else {
            return []int{i, j}
        }
    }
    return []int{}
}
