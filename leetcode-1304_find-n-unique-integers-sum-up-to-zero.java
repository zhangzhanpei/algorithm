/**
 * 给定一个 int 数组元素数量，使得该数组元素之和为 0
 */
public int[] sumZero(int n) {
    int[] res = new int[n];
    for (int i = 0; i < n - 1; i += 2) {
        res[i] = i + 1;
        res[i + 1] = -res[i];
    }
    return res;
}
