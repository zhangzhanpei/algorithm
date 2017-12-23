/**
 * 给定一个数组, 将它两两分组, 取小的那个元素求和使得和最大
 */
public static int arrayPairSum(int[] nums) {
    Arrays.sort(nums); //由小到大排序
    int sum = 0;
    for (int i = 0; i < nums.length; i = i + 2) { //两两分组取小的求和
        sum += nums[i];
    }
    return sum;
}
