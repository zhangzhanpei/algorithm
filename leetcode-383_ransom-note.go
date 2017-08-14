/**
 * 给定字符串 a 和 b, 判断 a 是否可以由 b 中的字符组成
 */
package main

import "fmt"

func main() {
    canConstruct("aa", "ab")
}

func canConstruct(ransomNote string, magazine string) bool {
    m := map[byte]int{}
    for i := 0; i < len(magazine); i++ {
        m[magazine[i]] += 1
    }

    for i := 0; i < len(ransomNote); i++ {
        if m[ransomNote[i]] - 1 < 0 {
            return false
        }
        m[ransomNote[i]] -= 1
    }
    return true
}
