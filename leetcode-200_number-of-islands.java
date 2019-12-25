/**
 * 给定一个由'0'和'1'组成的二维数组, '0'表示海水, '1'表示陆地, 求陆地的数量(相邻的'1'算一个陆地)
 */
public int numIslands(char[][] grid) {
    int num = 0;
    for (int i = 0; i < grid.length; i++) {
        for (int j = 0; j < grid[0].length; j++) {
            // 遇到一块陆地
            if (grid[i][j] == '1') {
                num++;
                // 递归地将当前陆地与相邻陆地变成海水(即整块陆地记为一块)
                dfs(grid, i, j);
            }
        }
    }
    return num;
}

public void dfs(char[][] grid, int i, int j) {
    // 如果超出边界或已经是海水，无需处理
    if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == '0') {
        return;
    }
    // 变成海水
    grid[i][j] = '0';
    // 继续处理相邻区域
    dfs(grid, i + 1, j);
    dfs(grid, i - 1, j);
    dfs(grid, i, j + 1);
    dfs(grid, i, j - 1);
}
