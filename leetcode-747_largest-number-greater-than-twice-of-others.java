/**
 * 给定一个int数组，找到最大元素的下标，最大元素起码是其他元素的两倍大
 */
public static int dominantIndex(int[] nums) {
    int maxOne = 0; //记录最大元素
    int maxTwo = 0; //次大元素
    int idx = 0;
    for (int i = 0; i < nums.length; i++) { //遍历一趟找到最大和次大的元素
        if (nums[i] > maxTwo) {
            maxTwo = nums[i];
            if (maxTwo > maxOne) {
                int tmp = maxOne;
                maxOne = maxTwo;
                maxTwo = tmp;
                idx = i;
            }
        }
    }
    return maxOne >= maxTwo * 2 ? idx : -1; //返回
}
