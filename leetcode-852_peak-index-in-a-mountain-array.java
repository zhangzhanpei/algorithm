/**
 * 给定一个 int 数组，返回一个峰顶元素的下标
 * 峰顶元素即 A[i - 1] < A[i] > A[i + 1] 中的 A[i]
 */
public int peakIndexInMountainArray(int[] A) {
    for (int i = 1; i < A.length; i++) {
        if (A[i] < A[i - 1]) {
            return i - 1;
        }
    }
    // 如果元素一直在增大，则最后一个元素即峰顶
    return A.length - 1;
}
