/**
 * 给定一个整数x, 求平方根
 * 牛顿迭代法: 先猜测平方根为y. 如果y不对或精度不够, 令y=(y+x/y)/2, 可无限逼近正确值
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
    e, y := 1e-5, 1.0
    for math.Abs(float64(x) - y * y) > e {
        y = (y + float64(x) / y) / 2
    }
    return int(y)
}
