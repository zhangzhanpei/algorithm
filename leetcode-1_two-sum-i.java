import java.util.Arrays;

/**
 * 给定一个有序 int 数组和一个 target, 从数组中取两个元素使得和等于 target，返回两个元素的下标
 */
class Solution {
    public static void main(String[] args) {
        int[] nums = {2, 7, 11, 15};
        System.out.println(Arrays.toString(twoSum(nums, 9)));;
    }

    public static int[] twoSum(int[] nums, int target) {
        int i = 0;
        int j = nums.length - 1;
        // 因为数组是有序的，从数组的头部和尾部开始加
        while (i < j) {
            int sum = nums[i] + nums[j];
            if (sum > target) { // 如果大了，说明相加的元素应减小，即右边指针前移
                j--;
            } else if (sum < target) { // 如果小了，说明相加的元素应加大，即左边指针后移
                i++;
            } else {
                return new int[]{i, j}; // 刚好相等
            }
        }
        return new int[0];
    }
}
