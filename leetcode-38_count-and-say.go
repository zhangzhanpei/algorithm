/**
 * 给定一个正整数 n, 返回相应的字符串
 * n = 1, return "1"
 * n = 2, return "11"
 * n = 3, return "21"
 * n = 4, return "1211" 这里的 n 是指第四次计算, 计算对象是上面的 21, 这里的计算结果为一个2和一个1, 即 1211
 * n = 5, return "111221"
 */
package main

import "fmt"
import "strconv"

func main() {
    countAndSay(1)
}

func countAndSay(n int) string {
    if n <= 0 {
        return ""
    }
    ret := "1"
    for n > 1 {
        tmp := ""
        for i := 0; i < len(ret); i++ {
            count := 1
            for i + 1 < len(ret) && ret[i] == ret[i + 1] {
                count++
                i++
            }
            tmp += strconv.Itoa(count) + string(ret[i]) //个数 + 元素
        }
        ret = tmp //将结果传入下一趟计算
        n--
    }
    return ret
}
