//给定一个int数组, 所有元素都需要出现两次, 找到只出现了一次的元素
package main

import "fmt"

func main() {
    singleNumber([]int{1, 2, 1, 3, 2, 5})
}

func singleNumber(nums []int) []int {
    m := map[int]bool{}
    res := []int{}
    for _, v := range nums {
        if m[v] { //之前出现过的元素直接删除键
            delete(m, v)
        } else {
            m[v] = true
        }
    }
    for k, _ := range m { //map里面剩下的就是出现一次的元素
        res = append(res, k)
    }
    return res
}
