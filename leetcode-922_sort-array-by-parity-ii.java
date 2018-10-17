/**
 * 给定一个 int 数组，其中奇数偶数各一半，对该数组进行排序使得偶数元素的下标为偶数，奇数元素下标为奇数
 */
class Solution {
    public int[] sortArrayByParityII(int[] A) {
        int[] ret = new int[A.length];
        int i = 0, j = 1;
        for (int k : A) {
            if (k % 2 == 0) {
                ret[i] = k;
                i += 2;
            } else {
                ret[j] = k;
                j += 2;
            }
        }
        return ret;
    }
}
