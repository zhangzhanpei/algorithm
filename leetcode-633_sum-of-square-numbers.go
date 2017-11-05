//给定一个整数n, 求 i*i + j*j = n
package main

import "fmt"
import "math"

func main() {
    judgeSquareSum(5);
}

func judgeSquareSum(c int) bool {
    i, j := 0, int(math.Sqrt(float64(c)))
    for i <= j {
        sum := i * i + j * j
        if sum > c {
            j--
        } else if sum < c {
            i++
        } else {
            return true
        }
    }
    return false
}
