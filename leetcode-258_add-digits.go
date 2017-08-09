/**
 * 给定一个非负整数num, 将每位数字相加, 如果和不是个位数, 则继续上述处理直到剩下个位数的和
 * 解一：循环处理
 * 解二：ret = 1 + (num - 1) % 9
 */
package main

import "fmt"

func main() {
    fmt.Println(addDigits(19))
}

//解一
func addDigits(num int) int {
    loop:
    sum := 0
    for num > 0 {
        sum += num % 10
        num = int(num / 10)
    }
    if sum >= 10 {
        num = sum
        goto loop
    }
    return sum
}

//解二
func addDigits(num int) int {
    return 1 + (num - 1) % 9
}
