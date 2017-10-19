/**
 * 给定一个无序int数组和一个target, 从数组中取两个元素使得和等于target
 */
package main

import "fmt"

func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for k, v := range nums {
        if _, ok := m[v]; ok {
            return []int{m[v], k}
        }
        //map存储target与当前元素的差值, 当迭代元素等于map中的差值, 说明之前有元素加上当前元素等于target
        m[target - v] = k
    }
    return []int{}
}
