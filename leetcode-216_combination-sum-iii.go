/**
 * 给定一个整数k和一个整数n, 使用k个1-9的数字之和等于n, 返回所有组合
 */
package main

import "fmt"

func main() {
    fmt.Println(combinationSum3(3, 15))
}

func combinationSum3(k int, n int) [][]int {
    ret := [][]int{}
    dfs(&ret, []int{}, k, 1, n)
    return ret
}

/**
 * 深度优先查找
 * @param  ret   *[][]int      保存所有组合
 * @param  tmp   []int         一个组合
 * @param  k     int           几个数
 * @param  start int           开始组合的数字
 * @param  n     int           差
 */
func dfs(ret *[][]int, tmp []int, k, start, n int) {
    //超过k个元素或当前元素和已大于n, 直接返回进入下一个元素尝试
    if len(tmp) > k || n < 0 {
        return
    }

    //如果刚好k个元素并且和等于n, 保存该组合
    if len(tmp) == k && n == 0 {
        oneAnswer := make([]int, k)
        //这里使用copy是因为tmp会继续进行下一个元素的尝试, 如果不用copy会影响到已保存到ret里的结果
        copy(oneAnswer, tmp)
        *ret = append(*ret, oneAnswer)
        return
    }

    //逐个元素尝试
    for i := start; i <= 9; i++ {
        tmp = append(tmp, i);
        dfs(ret, tmp, k, i + 1, n - i)
        //弹出最后一个元素, 使用下一个元素尝试
        tmp = tmp[:len(tmp) - 1]
    }
}
