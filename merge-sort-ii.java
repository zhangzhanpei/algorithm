import java.util.Arrays;
import java.util.LinkedList;

/**
 * 自底向上的归并排序
 */
class Solution {
    public static void main(String[] args) {
        int[] nums = {6, 2, 13, 10, 1, 0, 89, 24, 11, 100, 96};
        mergeSort(nums);
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

    public static void mergeSort(int[] nums) {
        int len = nums.length;
        int mergeSize = 1;
        while (mergeSize < len) {
            int left = 0, mid = 0, right = 0;

            // 如果未合并元素数量大于 mergeSize * 2，则可以进行两个 mergeSize 数量的数组进行合并
            while (left + mergeSize * 2 - 1 < len) {
                mid = left + mergeSize - 1;
                right = left + mergeSize * 2 - 1;
                mergeArray(nums, left, mid, right);
                left = right + 1;
            }

            // 如果未合并元素数量不够 mergeSize * 2，则左侧数组元素为 mergeSize 个，右侧数组用剩下的元素进行合并
            mid = left + mergeSize - 1;
            if (mid < len - 1) {
                mergeArray(nums, left, mid, len - 1);
            }

            // 合并范围扩大一倍
            mergeSize *= 2;
        }
    }
}
