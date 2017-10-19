/**
 * 输出杨辉三角的第k行
 */
package main

import "fmt"

func main() {
    getRow(3)
}

func getRow(rowIndex int) []int {
    ret := make([]int, rowIndex + 1)
    ret[0] = 1
    for i := 0; i <= rowIndex; i++ {
        for j := i; j > 0; j-- {
            /**
             * j循环一遍即求得第i层的元素
             * j = 0 [1 0 0 0]
             * j = 1 [1 1 0 0]
             * j = 2 [1 2 1 0]
             * j = 3 [1 3 3 1]
             */
            ret[j] = ret[j] + ret[j - 1]
        }
    }
    return ret
}
