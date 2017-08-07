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
            tmp := make([]int, len(sub)) //这里很奇怪, 直接用sub来append的话数据有问题
            copy(tmp, sub)
            tmp = append(tmp, val)
            res = append(res, tmp)
        }
    }
    return res
}
