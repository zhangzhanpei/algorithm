/**
 * 给定一个正整数, 返回相应的字符串, 规则如下
 * 1 -> "A"
 * 2 -> "B"
 * 3 -> "C"
 * ...
 * 26 -> "Z"
 * 27 -> "AA"
 * 28 -> "AB"
 */
package main

import "fmt"

func main() {
    fmt.Println(convertToTitle(52))
}

func convertToTitle(n int) string {
    res := ""
    for n > 0 {
        n--
        res = string('A' + n % 26) + res
        n = int(n / 26)
    }
    return res
}
