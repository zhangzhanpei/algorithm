/**
 * 给定一个出席记录如"LALL", A表示缺席, P表示出席, L表示迟到
 * 如果出席记录中缺席次数<=1次 或 没有连续3次的迟到, 返回true
 */
package main

import "fmt"
import "strings"

func main() {
    fmt.Println(checkRecord("LALL"))
}

func checkRecord(s string) bool {
    return !strings.Contains(s, "LLL") && strings.Count(s, "A") <= 1
}
