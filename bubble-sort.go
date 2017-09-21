/**
 * 冒泡排序
 * 逐个遍历元素, 如果当前元素比后一个元素小, 两个元素交换
 */
package main

import "fmt"

func main() {
    fmt.Println(bubbleSort([]int{2, 7, 12, 6, 0}))
}

func bubbleSort(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    return nums
}
