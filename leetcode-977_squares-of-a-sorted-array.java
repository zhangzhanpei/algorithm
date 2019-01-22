/**
 * 给定一个有序int数组(带负数)，返回每个元素的平方值仍保持有序
 * 使用两个指针从两头向中间计算，大的放后面
 */
public int[] sortedSquares(int[] A) {
    int len = A.length;
    int i = 0;
    int j = len - 1;
    int k = len - 1;
    int[] squares = new int[len];
    while (i <= j) {
        if (A[i] * A[i] <= A[j] * A[j]) {
            squares[k] = A[j] * A[j];
            j--;
        } else {
            squares[k] = A[i] * A[i];
            i++;
        }
        k--;
    }
    return squares;
}
