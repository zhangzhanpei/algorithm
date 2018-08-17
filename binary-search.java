/**
 * 二分查找, 在有序数组中查找目标元素
 */
public int search(int[] nums, int target) {
    int left = 0;
    int right = nums.length - 1;
    while (left <= right) {
        int mid = (left + right) / 2;
        if (target > nums[mid]) { // 目标元素比中点元素大，往右侧查找
            left = mid + 1;
        } else if (target < nums[mid]) { // 目标元素比中点元素小，往左侧查找
            right = mid - 1;
        } else { // 相等判断放最后是因为目标元素刚好等于中点元素的几率很小，尽快往左或往右查找
            return mid;
        }
    }
    return -1;
}
