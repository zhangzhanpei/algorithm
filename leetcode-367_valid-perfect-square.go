//给定一个整数, 判断它是否为某个数的平方
package main

import "fmt"

func main() {
    isPerfectSquare(16)
}

func isPerfectSquare(num int) bool {
    x := num
    for x * x > num { //牛顿开方逐步逼近平方根
        x = (x + num / x) / 2
    }
    return x * x == num
}
