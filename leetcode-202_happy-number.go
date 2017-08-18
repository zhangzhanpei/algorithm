/**
 * 给定一个整数n, 按位平方求和, 如果最后的和等于1则该数为 happy number, 返回 true
 * 19 即为 happy number
 * 1 * 1 + 9 * 9 = 82
 * 8 * 8 + 2 * 2 = 68
 * 6 * 6 + 8 * 8 = 100
 * 1 * 1 + 0 * 0 + 0 * 0 = 1
 */
package main

import "fmt"

func main() {
    fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
    m := map[int]bool{} //如果 sum 已计算过, 则直接返回 false 以免死循环计算
    loop:
        sum := 0
        for n > 0 {
            tmp := n % 10
            sum += tmp * tmp
            n /= 10
        }
        if sum == 1 {
            return true
        }
        if m[sum] {
            return false
        } else {
            m[sum] = true
            n = sum
            goto loop
        }
}
