//给定一个字符串, 判断字符串是否由一个字串重复多次
package main

import "fmt"
import "strings"

func main() {
    repeatedSubstringPattern("abab")
}

func repeatedSubstringPattern(s string) bool {
    if s == "" {
        return false
    }
    ss := strings.Repeat(s, 2)
    ss = ss[1:len(ss) - 1]
    return strings.Contains(ss, s)
}
