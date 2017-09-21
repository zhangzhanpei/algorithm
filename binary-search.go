/**
 * 二分查找, 在有序数组中查找目标元素
 */
package main

import "fmt"

func main() {
    fmt.Println(binarySearch([]int{1,2,6,8,12,15,23}, 15))
}

func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := (left + right) / 2
        if target > nums[mid] { //目标元素大于中点值, 那么目标元素在中点右侧, 继续查找右边元素
            left = mid + 1
        } else if target < nums[mid] { //目标元素小于中点值, 那么目标元素在中点左侧, 继续查找左边元素
            right = mid - 1
        } else {
            return mid
        }
    }
    return -1
}
