/**
 * 给定一个二进制字符串, 求01数量相等的子字符串数量(01都是连续的)
 */
package main

import "fmt"

func main() {
    fmt.Println(countBinarySubstrings("00110011"))
}

func countBinarySubstrings(s string) int {
    prevLen, curLen, ret := 0, 1, 0
    for i := 1; i < len(s); i++ {
        if s[i] == s[i - 1] {
            curLen++
        } else {
            prevLen = curLen
            curLen = 1
        }
        if prevLen >= curLen {
            ret++
        }
    }
    return ret
}
