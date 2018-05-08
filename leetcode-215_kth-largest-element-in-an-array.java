/**
 * 查找数组中第 k 大的元素
 */
public int findKthLargest(int[] nums, int k) {
    int left = 0, right = nums.length - 1;
    while (true) {
        int pos = partition(nums, left, right);
        // 如果相遇处坐标刚好是 k - 1 说明该位置的元素就是第 k 大的元素
        if (pos == k - 1) {
            return nums[pos];
        }
        // 如果 k - 1 位于左边说明第 k 大的元素在左边
        if (pos > k - 1) {
            right = pos - 1;
        } else { // 否则在右边分区中继续找
            left = pos + 1;
        }
    }
}

// 将数组划分为两部分，左侧为大于基数的，右侧为小于基数的元素
public int partition(int[] nums, int left, int right) {
    int base = nums[left];
    int i = left;
    int j = right;
    while (i < j) {
        // 从右向左找到第一个大于基数的元素
        while (nums[j] <= base && i < j) {
            j--;
        }
        // 从左往右找到第一个小于基数的元素
        while (nums[i] >= base && i < j) {
            i++;
        }
        // 如果 i 和 j 没有相遇则交换元素
        if (i < j) {
            int tmp = nums[i];
            nums[i] = nums[j];
            nums[j] = tmp;
        }
    }
    // 将 i 和 j 相遇处的元素与基数交换
    nums[left] = nums[i];
    nums[i] = base;
    // 返回 i 和 j 相遇处的下标
    return i;
}
