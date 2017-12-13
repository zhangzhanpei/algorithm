/**
 * 给定一个三角形的二维列表，求经过的路径元素之和最小
 * 动态规划递推公式 min[i][j] = triangle[i][j] + Math.min(min[i + 1][j], min[i + 1][j + 1])
 * 自底向上逐层计算即可
 */
public int minimumTotal(List<List<Integer>> triangle) {
    if (triangle.size() == 0) {
        return 0;
    }
    if (triangle.size() == 1) {
        return triangle.get(0).get(0);
    }
    int row = triangle.size();
    int col = triangle.get(row - 1).size();
    int[][] ret = new int[row][col]; //使用二维数组保存计算结果
    for (int i = 0; i < col; i++) {
        ret[row - 1][i] = triangle.get(row - 1).get(i); //最后一行不用计算
    }
    for (int i = row - 2; i >= 0; i--) { //从倒数第二行开始
        for (int j = 0; j <= i; j++) {
            ret[i][j] = triangle.get(i).get(j) + Math.min(ret[i + 1][j], ret[i + 1][j + 1]);
        }
    }
    return ret[0][0];
}
