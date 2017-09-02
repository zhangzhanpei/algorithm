/**
 * 给定n块砖头, 垒成楼梯状, 求能垒几层
 * 第一层1
 * 第二层1 1
 * 第三层1 1 1
 * ...
 * 使用总数n减去每层需要的砖头数量, 直到n为负数(剩余砖头不够垒一层了)
 */
package main

import "fmt"

func main() {
    fmt.Println(arrangeCoins(8))
}

func arrangeCoins(n int) int {
    var i = 0
    for n > 0 {
        i++
        n -= i //减去每层的砖头
    }
    if n == 0 { //刚好用完砖头
        return i
    } else { //不够下一层了
        return i - 1
    }
}
