/**
 * 给定一个仅包含0和1的int数组, 0, 10, 11都可以表示一个字符, 求最后一个字符是否单字符(单独的0)
 */
package main

import "fmt"

func main() {
    fmt.Println(isOneBitCharacter([]int{1,0,0}))
}

func isOneBitCharacter(bits []int) bool {
    i, n := 0, len(bits) //i表示字符的起始位置
    for i < n - 1 {
        if bits[i] == 0 { //如果是0仅能表示一个字符
            i++
        } else { //如果是1必须两位才能表示一个字符
            i += 2
        }
    }
    return i == n - 1 //如果最后一个元素是字符的开始, 即最后一个元素为单字符
}
