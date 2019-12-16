import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * 给定一个无序 int 数组和一个 target, 从数组中取两个元素使得和等于 target，返回两个元素的下标
 */
class Solution {
    public static void main(String[] args) {
        int[] nums = {2, 7, 11, 15};
        System.out.println(Arrays.toString(twoSum(nums, 9)));;
    }

    public static int[] twoSum(int[] nums, int target) {
        // 使用 map 存储 target 和数组元素的差值和下标(差值为键，下标为值)
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            // 如果当前元素刚好是 map 中的键，说明之前有一个元素与当前元素相加刚好等于 target
            if (map.containsKey(nums[i])) {
                return new int[]{map.get(nums[i]), i};
            }
            map.put(target - nums[i], i);
        }
        return new int[0];
    }
}
