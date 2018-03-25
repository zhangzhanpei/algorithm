/**
 * 给定一个 int 数组, 所有元素出现了两次只有一个元素只出现了一次, 找到这个元素
 */
public int singleNonDuplicate(int[] nums) {
    int left = 0, right = nums.length - 1;
    while (left < right) {
        int mid = left + (right - left) / 2;
        //如果中间点的数与左右邻居不相等
        if (nums[mid] != nums[mid + 1] && nums[mid] != nums[mid - 1]) return nums[mid];
        if (nums[mid] == nums[mid + 1] && mid % 2 == 0) { //mid 为偶数说明左侧的元素数量为偶数，如果 nums[mid] == nums[mid + 1] 说明单独数在 mid + 1 的右边
            left = mid + 1;
        } else if (nums[mid] == nums[mid - 1] && mid % 2 == 1) { //mid 为奇数说明左侧元素数量为奇数，如果 nums[mid] == nums[mid - 1] 说明 mid - 1 左侧的元素都有重复
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return nums[left];
}
