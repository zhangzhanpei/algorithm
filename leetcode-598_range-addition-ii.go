/**
 * 给定一个 m*n 的二维矩阵和一个操作数组, 求经过操作数组操作后二维矩阵最大元素的个数
 * 操作数组元素为坐标[a, b], 则[0, 0]到[a, b]的元素都自增1
 */
package main

import "fmt"
import "math"

func main() {
    fmt.Println(maxCount(3, 3, [][]int{
        {2, 2},
        {3, 3},
    }))
}

func maxCount(m int, n int, ops [][]int) int {
    //没有操作时, 整个矩阵的所有元素都是最大元素0
    if len(ops) == 0 {
        return m * n
    }

    //逐个取出操作
    for _, v := range ops {
        //最大的元素就是所有操作范围的交集(因为交集范围的元素自增次数最多), 也就是[0, 0]到[minM, minN]的元素
        m = int(math.Min(float64(m), float64(v[0])))
        n = int(math.Min(float64(n), float64(v[0])))
    }
    return m * n
}
