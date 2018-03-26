/**
 * 给定一数组有 1 ≤ a[i] ≤ n, 其中 n 是数组的长度
 * 数组中有一些数字重复出现了, 找出 1 ~ n 未在数组中出现的数字
 */
public List<Integer> findDisappearedNumbers(int[] nums) {
    List<Integer> ret = new ArrayList<>();
    //将元素对应位置的元素置为负数
    for (int i = 0; i < nums.length; i++) {
        int idx = Math.abs(nums[i]) - 1;
        nums[idx] = -Math.abs(nums[idx]);
    }
    //找到未被置为负数的下标，即未出现的数
    for (int j = 0; j < nums.length; j++) {
        if (nums[j] > 0) {
            ret.add(j + 1);
        }
    }
    return ret;
}
