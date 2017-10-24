/**
 * 给定一个int数组, 生成所有的不重复子数组
 */
package main

import "fmt"
import "sort"

func main() {
    fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

func subsets(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{{}} //空数组是任何数组的子数组
    for _, val := range nums {
        for _, sub := range res {
            tmp := make([]int, len(sub))
            //这里使用copy生成一个新的tmp用于res的append, 如果是直接append(res, sub)的话sub进入下一个迭代时会影响到ret结果
            copy(tmp, sub)
            tmp = append(tmp, val)
            res = append(res, tmp)
        }
    }
    return res
}
