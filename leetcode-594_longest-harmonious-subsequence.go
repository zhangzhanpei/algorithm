/**
 * 给定一个int数组, 返回最长和谐子序列的长度
 * 和谐指数组中最大值和最小值的差值为1, 即求数组中元素n和元素n+1的最大个数
 * 使用map记录各元素的个数, 相邻元素的个数之和取最大的即可
 */
package main

import "fmt"

func main() {
    fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
    //使用map记录各元素的个数
    m := map[int]int{}
    for _, val := range nums {
        m[val]++
    }

    max := 0
    for key, _ := range m {
        if m[key + 1] > 0 { //如果有下一个相邻元素
            tmp := m[key] + m[key + 1] //相邻元素个数相加
            if tmp > max { //取最大值
                max = tmp
            }
        }
    }

    return max
}
