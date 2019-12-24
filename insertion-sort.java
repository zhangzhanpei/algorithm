/**
 * 插入排序
 * 逐个遍历元素, 如果当前元素比前面的元素小, 一直向前交换直到位置正确
 * 即前面有序部分比当前元素大的元素都后移一位, 把当前元素交换(插入)到正确位置
 */
public int[] insertionSort(int[] nums) {
    for (int i = 1; i < nums.length; i++) {
        for (int j = i; j >= 1 && nums[j] < nums[j - 1]; j--) { // 如果当前元素比前面的小, 一直向前交换直到位置正确
            int tmp = nums[j - 1];
            nums[j - 1] = nums[j];
            nums[j] = tmp;
        }
    }
    return nums;
}
