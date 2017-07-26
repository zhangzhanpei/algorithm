package main

import "fmt"

var arr = [][]int{
    {7},
    {3, 8},
    {8, 1, 0},
    {2, 7, 4, 4},
    {4, 5, 2, 6, 5},
}
var book [5][5]int

func main() {
    fmt.Println(maxSum(0, 0))
}

//使用 book 数组保存递归过程中的计算结果
func maxSum(i, j int) int {
    if book[i][j] != 0 { //如果已计算过直接返回结果
        return book[i][j]
    }
    if i == 4 {
        book[i][j] = arr[i][j]
    } else {
        x := maxSum(i + 1, j) //走左下时最大值
        y := maxSum(i + 1, j + 1) //走右下
        sum := 0
        if x > y {
            sum = x + arr[i][j]
        } else {
            sum = y + arr[i][j]
        }
        book[i][j] = sum //保存递归过程的计算结果
    }
    return book[i][j]
}
