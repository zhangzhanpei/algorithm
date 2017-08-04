/**
 * 给定一个字符串, 出现频率最高的字符排前面
 */
package main

import "fmt"

func main() {
    fmt.Println(frequencySort("tree")) //output "eert"
}

func frequencySort(s string) string {
    //先把字符和出现的次数写进一个map
    m := map[byte]int{}
    n := len(s)
    for i := 0; i < n; i++ {
        m[s[i]] += 1
    }

    //次数为key, 字符拼接为值, 这里使用了桶排序
    bucket := make([]string, n + 1)
    for key, val := range m {
        for i := 0; i < val; i++ {
            bucket[val] += string(key)
        }
    }

    //从bucket后面往回读, 拼接字符串
    res := ""
    for i := len(bucket) - 1; i >= 0; i-- {
        if bucket[i] != "" {
            res += bucket[i]
        }
    }
    return res
}
