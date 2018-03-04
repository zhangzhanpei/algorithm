/**
 * 给定n个坐标, 找出回旋镖的个数
 * 回旋镖为3个坐标[a, b, c]组成, 并且 ab 的距离和 ac 的距离相同
 */
package main

import "fmt"

func main() {
    fmt.Println(numberOfBoomerangs([][]int{
        {0, 0},
        {1, 0},
        {2, 0},
    }))
}

func numberOfBoomerangs(points [][]int) int {
    ret := 0
    //逐个坐标作为起点, 求出其他坐标与它的距离
    for i := 0; i < len(points); i++ {
        m := map[int]int{}
        for j := 0; j < len(points); j++ {
            //勾股定理求斜边长度
            x := points[i][0] - points[j][0]
            y := points[i][1] - points[j][1]
            m[x * x + y * y]++
        }

        //当有n个坐标与起始坐标距离相等, 计算排列个数(从n个坐标中取出两个来组合)
        //A(n, 2) = n! / (n - 2)! = n * (n - 1)
        for _, v := range m {
            ret += v * (v - 1)
        }
    }
    return ret
}
