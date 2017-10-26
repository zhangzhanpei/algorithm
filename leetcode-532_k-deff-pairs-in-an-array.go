/**
 * 给定一个int数组和差k, 数组中任意两个元素的差为k, 求组合的个数
 */
package main

import "fmt"

func main() {
    fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
}

func findPairs(nums []int, k int) int {
    if len(nums) == 0 || k < 0 {
        return 0
    }

    //将元素出现的次数存到map中
    m := map[int]int{}
    for _, val := range nums {
        m[val] += 1
    }

    res := 0
    for key, val := range m {
        if k == 0 { //如果差是0, 即相同元素相减, 求map中次数大于2的元素即可
            if val >= 2 {
                res++
            }
        } else { //如果差不为0, 则当前元素加上k, 看有没有这个key即可
            if m[key + k] > 0 {
                res++
            }
        }
    }
    return res
}
