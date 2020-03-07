import java.util.Arrays;

/**
 * 自顶向下的归并排序
 * 参考文章 https://blog.csdn.net/yusiguyuan/article/details/42806763
 */
class Solution {
    public static void main(String[] args) {
        int[] nums = {10, 30, 50, 60, 80, 20, 25, 55, 65, 73};
        mergeSort(nums, 0, nums.length - 1);
        System.out.println(Arrays.toString(nums));
    }

    // 反转数组
    public static void reverse(int[] nums, int left, int right) {
        while (left < right) {
            int tmp = nums[left];
            nums[left] = nums[right];
            nums[right] = tmp;
            left++;
            right--;
        }
    }

    // 交换两个数组子集的位置
    public static void exchangeSubArray(int[] nums, int left, int mid, int right) {
        reverse(nums, left, mid);
        reverse(nums, mid + 1, right);
        reverse(nums, left, right);
    }

    // 原地合并两个有序数组
    public static void inPlaceMergeArray(int[] nums, int left, int mid, int right) {
        int i = left, j = mid + 1;
        while (i < j && j <= right) {
            int step = 0;
            while (i < j && nums[i] <= nums[j]) {
                i++;
            }
            while (j <= right && nums[i] >= nums[j]) {
                j++;
                step++;
            }
            exchangeSubArray(nums, i, mid, j - 1);
            i += step;
            mid += step;
        }
    }

    // 递归调用 mergeSort 方法进行数组拆分，当 left = right 时说明拆到了最小的数组，只剩一个元素
    // 当无法继续往下递归拆分数组时，返回上一层调用，开始拆分上一层右侧的数组
    // 当 left = 0, right = 0 时，回到上一层 left = 0, right = 1 执行 mergeSort(arr, 1, 1) 拆右边，此时右边也拆完，则执行 mergeArray(arr, 0, 1, 1) 进行合并数组
    public static void mergeSort(int[] nums, int left, int right) {
        if (left < right) {
            int mid = (left + right) / 2;
            mergeSort(nums, left, mid);
            mergeSort(nums, mid + 1, right);
            inPlaceMergeArray(nums, left, mid, right);
        }
    }
}
