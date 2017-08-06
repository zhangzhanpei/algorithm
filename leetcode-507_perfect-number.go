/**
 * 给定一个正整数, 判断该整数是否等于它所有整除数(自身除外)的和
 * 例如 28 = 1 + 2 + 4 + 7 + 14
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(checkPerfectNumber(28))
}

func checkPerfectNumber(num int) bool {
    if num == 1 { //1只有自身可整除, 返回false
        return false
    }
    sum := 0

    //这里从2开始, 从1开始的话会加到自身
    //循环时最大达到平方根即可
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num % i == 0 {
            sum += i + num / i //sum加上除数和商
        }
    }
    if sum + 1 == num { //这里的和加上之前跳过的1
        return true
    }
    return false
}
