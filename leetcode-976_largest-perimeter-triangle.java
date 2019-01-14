/**
 * 给定一个int数组，任取三个元素组成一个三角形，返回最大周长
 */
class Solution {
    public int largestPerimeter(int[] A) {
        // 先对数组排序
        Arrays.sort(A);
        // 从后往前找，第一个符合三角形条件的三个元素即周长最大
        int idx = A.length - 3;
        while (idx >= 0) {
            // 两边之和大于第三边
            if (A[idx] + A[idx + 1] > A[idx + 2]) {
                return A[idx] + A[idx + 1] + A[idx + 2];
            }
            idx--;
        }
        return 0;
    }
}
