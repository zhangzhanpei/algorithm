/**
 * 给定一个由0和1组成的二维数组, 0表示海水, 1表示陆地, 求陆地的数量(连续的1算一个陆地)
 */
func numIslands(grid [][]byte) int {
    num := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if string(grid[i][j]) == "1" {
                //如果当前坐标为陆地, 则将邻接当前坐标的陆地坐标都置0
                num++
                searchIsland(grid, i, j)
            }
        }
    }
    return num
}

func searchIsland(grid [][]byte, i int, j int) {
    //未超出边界并且为1则算陆地
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && string(grid[i][j]) == "1" {
        grid[i][j] = '0' //将当前坐标的上下左右邻接陆地坐标都置0, 当作一整块陆地
        searchIsland(grid, i + 1, j)
        searchIsland(grid, i - 1, j)
        searchIsland(grid, i, j + 1)
        searchIsland(grid, i, j - 1)
    }
}
