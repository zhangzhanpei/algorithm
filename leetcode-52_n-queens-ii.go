/**
 * 输出N皇后的解的数量
 */
var ret int

func totalNQueens(n int) int {
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
        ret++
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
