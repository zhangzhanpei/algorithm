package main

import "fmt"

var arr = [][]int{
    {7},
    {3, 8},
    {8, 1, 0},
    {2, 7, 4, 4},
    {4, 5, 2, 6, 5},
}
func main() {
    fmt.Println(maxSum(0, 0))
}

//直接递归解决
func maxSum(i, j int) int {
    if i == 4 {
        return arr[i][j]
    }
    x := maxSum(i + 1, j) //走左下时最大值
    y := maxSum(i + 1, j + 1) //走右下
    if x > y {
        return x + arr[i][j]
    } else {
        return y + arr[i][j]
    }
}
