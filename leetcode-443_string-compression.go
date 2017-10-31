/**
 * 给定一个byte数组, 原地压缩该数组
 * 如['a', 'a', 'b', 'b', 'b', 'c'] => ['a', '2', 'b', '3'. 'c']
 */
package main

import "fmt"
import "strconv"

func main() {
    fmt.Println(compress([]byte{'G','t','W','Y','v','&',':','d','#','k'}))
}

func compress(chars []byte) int {
    i, j, k, n := 0, 0, 0, len(chars) //i, j, k 三个下标
    for i < n {
        for j < n && chars[i] == chars[j] {
            j++ //j向后走到第一个不同字符的位置
        }
        chars[k] = chars[i] //保存相同字符为一个字符
        k++
        if j - i > 1 { //如果该字符的长度不为1
            tmp := []byte(strconv.Itoa(j - i)) //将长度转为[]byte格式跟在该字符后面
            for _, v := range tmp {
                chars[k] = v
                k++
            }
        }
        i = j //处理下一个不同字符
    }
    return k
}
