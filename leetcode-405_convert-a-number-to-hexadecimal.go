//10进制转16进制
package main

import "fmt"
import "math"

func main() {
    fmt.Println(toHex(-1))
}

func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    if num < 0 {
        num += math.MaxUint32 + 1
    }
    m := [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
    res := ""
    for num > 0 {
        res = string(m[num % 16]) + res
        num = num / 16
    }
    return res
}
