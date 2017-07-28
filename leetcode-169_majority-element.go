//给定一个数组, 其中有一个元素超过一半的数量, 找出该元素
package main

import "fmt"

func main() {
    fmt.Println(majorityElement([]int{6, 5, 5}))
}

func majorityElement(nums []int) int {
    m := map[int]int{} //定义一个map保存元素出现的次数
    for i := 0; i < len(nums); i++ {
        m[nums[i]] += 1
        if (m[nums[i]] > len(nums) / 2) {
            return nums[i]
        }
    }
    return 0
}
