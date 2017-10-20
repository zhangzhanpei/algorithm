/**
 * 给定一个二维数组存储图片的亮度, 对该图片进行平滑处理
 * 逐个坐标加上周围坐标的亮度取均值即可
 */
package main

import "fmt"

func main() {
    imageSmoother([][]int{
        {1, 1, 1},
        {1, 0, 1},
        {1, 1, 1},
    })
}

func imageSmoother(M [][]int) [][]int {
    if len(M) == 0 || len(M[0]) == 0 {
        return [][]int{}
    }
    m, n := len(M), len(M[0])
    dirs := [][]int{
        {-1, 1}, //左上
        {0, 1}, //正上
        {1, 1}, //右上
        {-1, 0}, //左
        {1, 0}, //右
        {-1, -1}, //左下
        {0, -1}, //正下
        {1, -1}, //右下
    }
    ret := [][]int{}
    for i := 0; i < m; i++ {
        tmp := make([]int, n)
        for j := 0; j < n; j++ {
            cur := M[i][j]
            count := 1
            for _, val := range dirs { //每个坐标加上周围8个坐标的元素求平均值
                x, y := i + val[0], j + val[1]
                if x < 0 || x >= m || y < 0 || y >= n { //超出范围的坐标不作计算
                    continue
                }
                count++
                cur += M[x][y]
            }
            tmp[j] = cur / count
        }
        ret = append(ret, tmp)
    }
    return ret
}
