/**
 * 给定两个int数组求交集(元素可重复)
 */
package main

import "fmt"

func main() {
    fmt.Println(intersect([]int{1,2,2,1}, []int{2,2}))
}

func intersect(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    for _, v := range nums1 { //记录数组一中的元素出现次数
        m[v]++
    }

    ret := []int{}
    for _, v := range nums2 { //从数组二中取回相应次数的相应元素
        if m[v] > 0 {
            ret = append(ret, v)
            m[v]--
        }
    }
    return ret
}
