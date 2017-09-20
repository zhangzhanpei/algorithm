/**
 * N皇后问题
 */
package main

import "fmt"

func main() {
    fmt.Println(solveNQueens(5))
}

var ret [][]string

func solveNQueens(n int) [][]string {
    nQueens := [][]string{}
    for i := 0; i < n; i++ { //初始化一个二维数组
        nQueens = append(nQueens, make([]string, n))
    }

    dfs(nQueens, 0, n)
    return ret
}

//递归求N皇后解法
func dfs(nQueens [][]string, row, n int) {
    if row == n { //到达最后一行
        //nQueens数组即一个解, 这里拼成字符串是题目要求
        tmp := []string{}
        for i := 0; i < n; i++ {
            str := ""
            for j := 0; j < n; j++ {
                if nQueens[i][j] != "" {
                    str += nQueens[i][j]
                } else {
                    str += "."
                }
            }
            tmp = append(tmp, str)
        }
        ret = append(ret, tmp)
        return
    }

    //在新行内逐个列尝试
    for col := 0; col < n; col++ {
        if isValid(nQueens, row, col, n) {
            nQueens[row][col] = "Q" //放置皇后
            dfs(nQueens, row + 1, n) //进入下一行
            nQueens[row][col] = "." //取出该位置的皇后, 尝试下一个col
        }
    }
}

//判断某个坐标能否放置皇后
func isValid(nQueens [][]string, row, col, n int) bool {
    //判断已有皇后的行, 当前列是否有冲突
    for i := 0; i < row; i++ {
        if nQueens[i][col] == "Q" {
            return false
        }
    }

    //判断左上方斜线上是否有皇后
    for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
        if nQueens[i][j] == "Q" {
            return false
        }
    }

    //判断右上方斜线上是否有皇后
    for i, j := row - 1, col + 1; i >= 0 && j < n; i, j = i - 1, j + 1 {
        if nQueens[i][j] == "Q" {
            return false
        }
    }

    return true
}
