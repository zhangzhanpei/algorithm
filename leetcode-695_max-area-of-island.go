/**
 * 给定一个由0和1组成的二维数组, 0表示海水, 1表示陆地, 求陆地最大面积(即连续的1的个数)
 */
func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                //逐个坐标为中心求周围连续的1
                maxArea = int(math.Max(float64(maxArea), float64(areaOfIsland(grid, i, j))))
            }
        }
    }
    return maxArea
}

func areaOfIsland(grid [][]int, i int, j int) int {
    //未超出边界并且为1则参与计算
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
        grid[i][j] = 0 //参与过计算的坐标置0
        return 1 + areaOfIsland(grid, i + 1, j) + areaOfIsland(grid, i - 1, j) + areaOfIsland(grid, i, j + 1) + areaOfIsland(grid, i, j - 1)
    }
    return 0
}
