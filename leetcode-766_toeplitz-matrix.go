/**
 * 给定一个二维int数组，确定是否toeplitz-matrix
 * toeplitz-matrix：矩阵的左上角到右下角的对角线上的元素都相等
 */
func isToeplitzMatrix(matrix [][]int) bool {
    for i := 0; i < len(matrix) - 1; i++ {
        for j := 0; j < len(matrix[0]) - 1; j++ {
            if matrix[i][j] != matrix[i + 1][j + 1] {
                return false
            }
        }
    }
    return true
}
