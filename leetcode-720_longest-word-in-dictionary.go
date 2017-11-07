/**
 * 给定一个string数组, 找到最长的单词
 * 最长单词由数组中其它单词一次组成, 当两个单词长度一致时, 返回字典序前面的
 */
package main

import "fmt"
import "sort"

func main() {
    fmt.Println(longestWord([]string{"w","wo","wor","worl", "world"}))
}

func longestWord(words []string) string {
    sort.Strings(words) //string数组按字典序排序
    m := map[string]bool{}
    ret := ""
    for _, w := range words {
        _, ok := m[w[:len(w) - 1]]
        if ok || len(w) == 1 { //如果单词长度为1或者可由前面的单词加一个字符组成
            if len(w) > len(ret) { //和之前的最长单词比较长度, 只有更长才替换, 这样可以保留字典序第一个最长单词
                ret = w
            }
            m[w] = true
        }
    }
    return ret
}
