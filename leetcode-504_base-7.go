/**
 * 给定一个整数, 返回它的7进制字符串
 */
package main

import "fmt"
import "math"
import "strconv"

func main() {
    fmt.Println(convertToBase7(0))
}

func convertToBase7(num int) string {
    n := int(math.Abs(float64(num)))
    s := ""
    for {
        s = strconv.Itoa(n % 7) + s
        n /= 7
        if n == 0 {
            break
        }
    }
    if num < 0 {
        s = "-" + s
    }
    return s
}
