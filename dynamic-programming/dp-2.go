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

//递归优化成递推
func main() {
    //先填充最后一行
    for i := 0; i < 5; i++ {
        book[4][i] = arr[4][i]
    }
    for x := 3; x >= 0; x-- { //从倒数第二行往上填充 book 数组
        for y := 0; y <= x; y++ {
            max := 0
            if book[x + 1][y] > book[x + 1][y + 1] { //取下一行左右两个元素中较大那个加上当前元素
                max = book[x + 1][y]
            } else {
                max = book[x + 1][y + 1]
            }
            book[x][y] = max + arr[x][y] //保存到 book 中
        }
    }
    fmt.Println(book[0][0]) //到第一行时就是最后的和
}
