/**
 * 给定两个int数组求交集
 */
package main

import "fmt"

func main() {
    fmt.Println(intersection([]int{1,2,2,1}, []int{2,2}))
}

func intersection(nums1 []int, nums2 []int) []int {
    m := map[int]bool{}
    for _, v := range nums1 { //第一个数组写进map
        m[v] = true
    }

    ret := []int{}
    for _, v := range nums2 { //遍历第二个数组中的元素，查看是否在map中
        if m[v] {
            ret = append(ret, v)
            delete(m, v)
        }
    }
    return ret
}
