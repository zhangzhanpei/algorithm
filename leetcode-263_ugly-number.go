/**
 * 给定一个整数 num, 判断是否 ugly number
 * 将给定数除以2、3、5，直到无法整除，也就是除以2、3、5的余数不再为0时停止。这时如果得到1，说明是所有因子都是2或3或5。
 * 如果不是1，则不是丑陋数。
 */
package main

import "fmt"

func main() {
    fmt.Println(isUgly(6))
}

func isUgly(num int) bool {
    for _, val := range []int{2, 3, 5} {
        for num > 0 && num % val == 0 {
            num /= val
        }
    }
    return num == 1
}
