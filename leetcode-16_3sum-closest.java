/**
 * 给定一个 int 数组和一个目标 target，其中存在三个元素组成的三元组 (a, b, c)，使得 a + b + c 最接近 target，返回这个三元组之和
 */
public int threeSumClosest(int[] nums, int target) {
    // 先随便取3个元素求和
    int ret = nums[0] + nums[1] + nums[2];
    Arrays.sort(nums);
    int sum, left, right;
    for (int i = 0; i < nums.length; i++) {
        left = i + 1;
        right = nums.length - 1;
        while (left < right) {
            sum = nums[i] + nums[left] + nums[right];
            // 如果和刚好等于target，直接返回
            if (sum == target) {
                return sum;
            } else if (sum > target) {
                right--;
            } else {
                left++;
            }
            // 看这个三元组之和与目标的绝对值，如果比之前的更小(更接近target)，更新为当前三元组之和
            if (Math.abs(sum - target) < Math.abs(ret - target)) {
                ret = sum;
            }
        }
    }
    return ret;
}
