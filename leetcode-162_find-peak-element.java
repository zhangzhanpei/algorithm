/**
 * 给定一个 int 数组，返回一个峰顶元素的下标
 * 峰顶元素即 nums[i - 1] < nums[i] > nums[i + 1] 中的 nums[i]
 */
public int findPeakElement(int[] nums) {
    for (int i = 1; i < nums.length; i++) {
        if (nums[i] < nums[i - 1]) {
            return i - 1;
        }
    }
    // 如果元素一直在增大，则最后一个元素即峰顶
    return nums.length - 1;
}
