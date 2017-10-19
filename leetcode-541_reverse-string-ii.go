//给定一个字符串和整数k, 反转所有2k位置的k个字符
package main

import "fmt"

func main() {
    fmt.Println(reverseStr("abcdefg", 2))
}

func reverseStr(s string, k int) string {
    char := []byte(s)
    n := len(char)
    for i := 0; i < n; {
        j := i + k - 1
        if j > n - 1 {
            j = n - 1
        }
        swap(char, i, j)
        i += 2 * k
    }
    return string(char)
}

func swap(arr []byte, i, j int) {
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}
