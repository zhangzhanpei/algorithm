//给定一个int数组, 将值为0的元素移动到后面, 其余元素保持原来的顺序
package main

import "fmt"

func main() {
    moveZeroes([]int{0, 1, 0, 3, 12});
}

func moveZeroes(nums []int)  {
    n := len(nums)
    for i := 0; i < n; i++ { //从前向后找到第一个0
        if nums[i] != 0 {
            continue
        }
        for j := i; j < n; j++ { //从0往后找到一个不为0的数交换
            if nums[j] != 0 {
                nums[i], nums[j] = nums[j], nums[i]
                break
            }
        }
    }
    return nums
}
