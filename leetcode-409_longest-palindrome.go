/**
 * 给定一个字符串, 取其中任意字符组成一个回文字符串, 求回文字符串的最大长度
 * 使用一个map, 字符第一次进map, 第二次出现则记两个长度
 */
package main

import "fmt"

func main() {
    longestPalindrome("abccccdd")
}

func longestPalindrome(s string) int {
    m := map[byte]int{}
    length := 0
    for i := 0; i < len(s); i++ {
        if m[s[i]] > 0 { //如果字符在前面出现过, 即可以和当前字符形成回文
            length += 2 //长度+2
            m[s[i]]--
        } else {
            m[s[i]]++
        }
    }
    for _, val := range m { //找出成对出现的字符后, 剩下的就是单独出现的字符
        if val > 0 { //如果还有单个的字符, 长度+1即可以组成一个奇数长度的回文字符串
            return length + 1
        }
    }
    return length
}
