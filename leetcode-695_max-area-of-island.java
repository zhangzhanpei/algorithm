/**
 * 给定一个由0和1组成的二维数组, 0表示海水, 1表示陆地, 求陆地最大面积(即连续的1的个数)
 */
public int maxAreaOfIsland(int[][] grid) {
    int maxArea = 0;
    for (int i = 0; i < grid.length; i++) {
        for (int j = 0; j < grid[0].length; j++) {
            // 遇到一块陆地
            if (grid[i][j] == 1) {
                maxArea = Math.max(maxArea, dfs(grid, i, j));
            }
        }
    }
    return maxArea;
}

// 递归地将当前陆地与相邻陆地变成海水(即整块陆地记为一块)
public int dfs(int[][] grid, int i, int j) {
    // 如果超出边界或已经是海水，无需处理
    if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == 0) {
        return 0;
    }
    // 变成海水
    grid[i][j] = 0;
    // 继续处理相邻区域
    return 1 + dfs(grid, i + 1, j) + dfs(grid, i - 1, j) + dfs(grid, i, j + 1) + dfs(grid, i, j - 1);
}
