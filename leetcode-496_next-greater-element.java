/**
 * 给定两个数组 nums1 和 nums2, nums1 是 nums2 的子集, 迭代 nums1 的元素, 在 nums2 中找出比该元素大的下一个元素位置, 找不到返回 -1
 */
public static int[] nextGreaterElement(int[] nums1, int[] nums2) {
    List<Integer> al = new ArrayList<>();
    for (int val : nums1) {
        //从出现位置往后找到第一个比val大的元素
        for (int j = indexOf(nums2, val); j <= nums2.length; j++) {
            if (j == nums2.length) {
                al.add(-1);
                break;
            }
            if (nums2[j] > val) {
                al.add(nums2[j]);
                break;
            }
        }
    }
    return al.stream().mapToInt(ele -> ele).toArray();
}

//查找元素的第一个出现位置
public static int indexOf(int[] nums, int target) {
    for (int i = 0; i < nums.length; i++) {
        if (nums[i] == target) return i;
    }
    return -1;
}
