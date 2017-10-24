/**
 * 给定一个int数组和一个目标值target, 求数组元素的和等于target的所有组合(可重复使用数组元素)
 */
package main

import "fmt"

func main() {
    fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
    ret := [][]int{}
    dfs(&ret, []int{}, candidates, 0, target)
    return ret
}

func dfs(ret *[][]int, tmp []int, nums []int, start int, n int) {
    if n < 0 {
        return
    }
    if n == 0 {
        oneAnswer := make([]int, len(tmp))
        copy(oneAnswer, tmp)
        *ret = append(*ret, oneAnswer)
        return
    }
    for i := start; i < len(nums); i++ {
        tmp = append(tmp, nums[i])
        dfs(ret, tmp, nums, i, n - nums[i]) //这里的i不用加1, 因为可以多次取同一个元素
        tmp = tmp[:len(tmp) - 1]
    }
}
