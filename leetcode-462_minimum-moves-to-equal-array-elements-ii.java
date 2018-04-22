/**
 * 给定一个int数组, 每一步可对其中一个元素+1或-1, 需要多少步能使所有元素相等
 */
public int minMoves2(int[] nums) {
    Arrays.sort(nums); //先排序求中间元素
    int mid = nums[nums.length / 2];
    int steps = 0;
    for (int val : nums) { //逐个元素计算需要多少步能变成中间元素的值, 将步数求和
        steps += Math.abs(val - mid);
    }
    return steps;
}
