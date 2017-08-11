//给定一正整数n, 求0-n每个数的二进制表示中1的个数
package main

import "fmt"

func main() {
    fmt.Println(0 >> 1)
}

func countBits(num int) []int {
    res := make([]int, num + 1)
    for i := 0; i <= num; i++ {
        res[i] = res[i >> 1] + (i & 1)
    }
    return res
}
