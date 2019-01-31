/**
 * 给定两个有序数组，求中值
 */
public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
    // 先合并两个数组
    List<Integer> l = new ArrayList<>();
    int i = 0, j = 0;
    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] <= nums2[j]) {
            l.add(nums1[i++]);
        } else {
            l.add(nums2[j++]);
        }
        System.out.println(666);
    }
    while (i < nums1.length) {
        l.add(nums1[i++]);
    }
    while (j < nums2.length) {
        l.add(nums2[j++]);
    }
    // 取中值
    if (l.size() % 2 == 0) {
        int mid = l.size() / 2;
        double d = l.get(mid - 1) + l.get(mid);
        return d / 2;
    } else {
        int mid = l.size() / 2;
        return l.get(mid);
    }
}
