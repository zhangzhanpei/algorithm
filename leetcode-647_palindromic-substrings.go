/**
 * 给定一个字符串, 求回文子字符串的数量
 */
package main

import "fmt"

func main() {
    fmt.Println(countSubstrings("aaa"))
}

var count int
func countSubstrings(s string) int {
    count = 0
    if s == "" {
        return 0
    }
    //将当前字符作为回文子字符串的中点向两边扩展
    for i := 0; i < len(s); i++ {
        extendPalindrome(s, i, i) //奇数长度的回文字符串, i 位置字符就是中间字符
        extendPalindrome(s, i, i + 1) //偶数长度的回文字符串, i 和 i + 1 位置字符就是中间的两个字符
    }
    return count
}

func extendPalindrome(s string, left int, right int) {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        count++
        left-- //向两边扩展
        right++
    }
}
