/**
 * 给定两个数字字符串, 不使用内置类型转换函数, 求两数的和
 */
package main

import "fmt"

func main() {
    fmt.Println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    var carry byte
    var ret string
    for i, j := m - 1, n - 1; i >= 0 || j >= 0; {
        var x, y byte
        if i >= 0 {
            x = num1[i] - '0' //逐个字符取到的是ascii, 减去0得到相应的数字
        } else {
            x = 0
        }
        if j >= 0 {
            y = num2[j] - '0'
        } else {
            y = 0
        }
        tmp := x + y + carry
        carry = tmp / 10
        ret = string(tmp % 10 + '0') + ret
        i--
        j--
    }

    if carry == byte(1) { //如果还有进位, 需要拼上进位的1
        ret = string(carry + '0') + ret
    }
    return ret
}
