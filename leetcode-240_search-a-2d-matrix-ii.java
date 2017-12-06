/**
 * 给定一个二维int数组，每个子数组都为升序，求给定目标值是否在二维数组中
 */
public boolean searchMatrix(int[][] matrix, int target) {
    if (matrix.length == 0 || matrix[0].length == 0) {
        return false;
    }
    int row = 0;
    int col = matrix[0].length - 1;
    while (row < matrix.length && col >= 0) { //从第0行开始
        if (matrix[row][col] == target) { //用每行的最后一个元素进行比较
            return true;
        } else if (matrix[row][col] < target) { //如果小于目标值则行增加
            row++;
        } else if (matrix[row][col] > target) { //如果大于目标值则列减少
            col--;
        }
    }
    return false;
}
