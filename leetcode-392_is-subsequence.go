/**
 * 给定两个字符串s和t, 判断s是否由t中的多个字符组成(从t中取的字符相对顺序不变)
 */
package main

import "fmt"

func main() {
    fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
    index := 0
    for i := 0; i < len(t) - 1; i++ { //s由左到右从t中取, 每取到一个则下标+1
        if s[index] != t[i] {
            continue
        }
        index += 1;
        if index == len(s) - 1 { //如果全部取到则返回true
            return true
        }
    }
    return false
}
