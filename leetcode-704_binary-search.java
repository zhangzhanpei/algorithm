/**
 * 二分查找, 在有序数组中查找目标元素
 */
public int search(int[] nums, int target) {
    int left = 0;
    int right = nums.length - 1;
    while (left <= right) {
        int mid = (left + right) / 2;
        if (target > nums[mid]) {
            left = mid + 1;
        } else if (target < nums[mid]) {
            right = mid - 1;
        } else {
            return mid;
        }
    }
    return -1;
}
