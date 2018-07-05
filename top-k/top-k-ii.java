/**
 * 快排划分方案
 */
public static void topK(int[] nums, int k) {
    partition(nums, 0, nums.length - 1, k);
}

public static void partition(int[] nums, int left, int right, int k) {
    if (left >= right) {
        return;
    }

    int pivot = nums[left];
    int i = left;
    int j = right;
    while (i < j) {
        // 从右往左找到第一个小于 pivot 的元素
        while (nums[j] >= pivot && i < j) {
            j--;
        }
        // 从左往右找到第一个大于 pivot 的元素
        while (nums[i] <= pivot && i < j) {
            i++;
        }
        // 两个元素在基数左右两边时交换位置
        if (i < j) {
            int tmp = nums[i];
            nums[i] = nums[j];
            nums[j] = tmp;
        }
    }
    // 一趟之后，比基数小的交换到了左边，比基数大的交换到了右边
    // 将基数交换到 i 和 j 相遇的位置
    nums[left] = nums[i];
    nums[i] = pivot;

    // 检查基数位置
    // 如果基数后面的元素数量不够 k 个，说明左边还有 top k 的元素，继续对左边元素进行划分
    if (i + k > nums.length) {
        partition(nums, left, i - 1, k);
    // 如果基数后面的元素数量多于 k 个，说明右边有不是 top k 的元素，继续对右边元素进行划分
    } else if (i + k < nums.length) {
        partition(nums, i + 1, right, k);
    // 如果右边刚好是 k 个元素，即 top k 元素
    } else {
        System.out.println(Arrays.toString(nums) + " " + i);
    }
}
