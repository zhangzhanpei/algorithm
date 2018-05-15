/**
 * 给定一个二维数组表示图片，翻转这张图片
 * 翻转：先按行反转数组，再将0改1，1改0
 */
public int[][] flipAndInvertImage(int[][] A) {
    for (int i = 0; i < A.length; i++) {
        int n = A[i].length;
        int[] tmp = new int[n];
        for (int j = 0; j < n; j++) {
            tmp[j] = A[i][n - j - 1] ^ 1;
        }
        A[i] = tmp;
    }
    return A;
}
