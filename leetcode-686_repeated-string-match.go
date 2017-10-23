/**
 * 给定一个短字符串A和长字符串B, 求A在几次重复后可以包含子串B
 */
package main

import "fmt"
import "math"
import "strings"

func main() {
    fmt.Println(repeatedStringMatch("a", "a"))
}

func repeatedStringMatch(A string, B string) int {
    lenA, lenB := len(A), len(B)
    tmp := ""
    //最少需要一次, 最多为拼到B的长度+1次
    for i := 1; i < int(math.Ceil(float64(lenB) / float64(lenA))) + 2; i++ {
        tmp += A
        if strings.Contains(tmp, B) {
            return i
        }
    }
    return -1
}
