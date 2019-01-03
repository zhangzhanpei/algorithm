/**
 * 给定一个int数组，判断数组是否为山峰状
 * 只有一个峰顶，两边元素都小于峰顶元素，峰顶不能是首元素或尾元素
 */
class Solution {
    public static boolean validMountainArray(int[] A) {
        if (A.length < 3) {
            return false;
        }
        int left = 0, right = A.length - 1;
        // 从左往右找到最大元素的下标
        while (left < A.length - 1 && A[left] < A[left + 1]) {
            left++;
        }
        // 从右往左找到最大元素的下标
        while (right > 0 && A[right] < A[right - 1]) {
            right--;
        }
        // 看两个下标是否同一个
        return left == right && left != 0 && right != A.length - 1;
    }
}
