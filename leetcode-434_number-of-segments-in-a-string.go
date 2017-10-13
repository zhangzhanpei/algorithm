/**
 * 给定一个 string, 判断它由多少个非空白字符段组成
 */
package main

import "fmt"
import "strings"

func main() {
    fmt.Println(countSegments("    "))
}

func countSegments(s string) int {
    strArr := strings.Split(s, " ") //先按空格切分字符串
    res := 0
    for _, val := range strArr { //判断非空白元素个数
        if val != "" {
            res += 1
        }
    }
    return res
}
