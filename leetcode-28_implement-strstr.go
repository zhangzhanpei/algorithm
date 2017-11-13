/**
 * 给定一个字符串和子串，返回子串在字符串中的位置或-1
 */
package main

import "fmt"

func main() {
    fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
    lh, ln := len(haystack), len(needle)
    for i := 0; i <= lh; i++ {
        for j := 0; j <= ln; j++ {
            if j == ln { //needle正确匹配到最后
                return i
            }
            if i + j >= lh { //haystack 比较完而 needle 未比较完
                return -1
            }
            if haystack[i + j] != needle[j] { //有字符不匹配
                break
            }
        }
    }
    return -1
}

// func strStr(haystack string, needle string) int {
//     return strings.Index(haystack, needle)
// }
