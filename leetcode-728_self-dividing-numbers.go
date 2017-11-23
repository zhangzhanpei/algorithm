/**
 * 给定一个int范围，返回该范围内的Self Dividing Numbers
 * Self Dividing Numbers : 该数可以被自身任意一位整除，且任意一位不为0
 */
package main

import "fmt"

func main() {
    fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
    ret := []int{}
    for i := left; i <= right; i++ {
        if isSelfDividing(i) {
            ret = append(ret, i)
        }
    }
    return ret
}

func isSelfDividing(num int) bool {
    n := num
    for n != 0 {
        tmp := n % 10
        if tmp == 0 { //不包含0
            return false
        }
        if num % tmp != 0 { //可被任意一位整除
            return false
        }
        n /= 10
    }
    return true
}
