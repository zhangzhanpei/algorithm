/**
 * 给定两个字符串 pattern 和 str, 判断他们是否对应关系
 * pattern = "abba", str = "dog cat cat dog" return true.
 * pattern = "abba", str = "dog cat cat fish" return false.
 */
package main

import "fmt"
import "strings"

func main() {
    fmt.Println(wordPattern("abba", "cat dog dog cat"))
}

func wordPattern(pattern string, str string) bool {
    strArr := strings.Split(str, " ")
    if len(pattern) != len(strArr) {
        return false
    }
    //使用map记录pattern字符和strArr字符串的对应关系
    m := map[rune]string{}
    for key, val := range pattern {
        if m[val] != "" && m[val] != strArr[key] {
            return false
        } else {
            m[val] = strArr[key]
        }
    }
    //当两个字符对应同一个字符串时, 返回false
    //如 pattern = "abba", str = "dog dog dog dog", a和b都是对应dog, return false.
    m1 := map[string]bool{}
    for _, val := range m {
        if m1[val] {
            return false
        } else {
            m1[val] = true
        }
    }
    return true
}
