/**
 * 给定一个 int 数组，判断该数组是否单调递增或递减的
 */
class Solution {
    public boolean isMonotonic(int[] A) {
        return this.asc(A) || this.desc(A);
    }

    // 单调递增
    public boolean asc(int[] A) {
        int n = A.length;
        for (int i = 0; i < n - 1; i++) {
            if (A[i] > A[i + 1]) {
                return false;
            }
        }
        return true;
    }

    // 单调递减
    public boolean desc(int[] A) {
        int n = A.length;
        for (int i = 0; i < n - 1; i++) {
            if (A[i] < A[i + 1]) {
                return false;
            }
        }
        return true;
    }
}
