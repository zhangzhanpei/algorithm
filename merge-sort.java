import java.util.Arrays;
import java.util.LinkedList;

/**
 * 自顶向下的归并排序
 */
class Solution {
    public static void main(String[] args) {
        int[] nums = {6, 2, 13, 10, 1, 0, 89, 24, 11, 100};
        mergeSort(nums, 0, nums.length - 1);
        System.out.println(Arrays.toString(nums));
    }

    // 合并两个有序数组
    public static void mergeArray(int[] nums, int left, int mid, int right) {
        int i = left;
        int j = mid + 1;
        LinkedList<Integer> tmp = new LinkedList<>();

        // 合并两个数组直到其中一个数组先取完
        while (i <= mid && j <= right) {
            if (nums[i] <= nums[j]) {
                tmp.add(nums[i]);
                i++;
            } else {
                tmp.add(nums[j]);
                j++;
            }
        }

        // 如果是右侧数组先取完，将左侧剩余元素追加到 tmp 中
        while (i <= mid) {
            tmp.add(nums[i]);
            i++;
        }

        // 如果是左侧数组先取完，将右侧剩余元素追加到 tmp 中
        while (j <= right) {
            tmp.add(nums[j]);
            j++;
        }

        // 将排好序的元素替换回原数组相应区间中
        for (int k = left; k <= right; k++) {
            nums[k] = tmp.removeFirst();
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
            mergeArray(nums, left, mid, right);
        }
    }
}
