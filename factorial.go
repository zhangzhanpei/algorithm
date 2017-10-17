/**
 * 给定一个整数n, 求n!
 * 注意 n > 12 时 int 的溢出问题
 */
package main

import "fmt"

func main() {
    factorial(361)
}

func factorial(n int) []int {
    ret := []int{1}
    //使用数组倒序存储阶乘结果
    for i := 2; i <= n; i++ {
        //进位
        carry := 0
        //乘以下一个数 i 时, 数组每一个元素都乘以 i
        for j := 0; j < len(ret); j++ {
            //前一个元素成以 i 的进位加到下一个元素乘积上
            tmp := ret[j] * i + carry
            carry = tmp / 10
            ret[j] = tmp % 10
        }

        //最后将进位按位跟到数组后面
        for carry > 0 {
            ret = append(ret, carry % 10)
            carry /= 10
        }
    }

    //倒序打印数组即 n 的阶乘结果
    for i := len(ret) - 1; i >= 0; i-- {
        fmt.Print(ret[i])
    }
    return ret
}
