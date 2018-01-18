/**
 * 给定一个由 0 和 1 组成的数组表示花床, 为 0 的位置可以种花, 但两棵花不能相邻. 给定 n 棵花, 求是否可以将这 n 棵花种到花床中
 */
package main

import "fmt"

func main() {
    fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    l := len(flowerbed)
    sli := make([]int, l + 2)
    copy(sli[1:], flowerbed) //flowerbed 两头加 0
    for i := 1; i < len(sli) - 1; i++ {
        if n == 0 {
            return true
        }
        if sli[i - 1] + sli[i] + sli[i + 1] == 0 { //如果前后都是 0 则可以种下
            i++ //这里和 for 循环中 i++ 共执行了两次, 即种下一棵花后跳过下一个位置
            n-- //每种下一棵花则 n--
        }
    }
    return n <= 0
}
