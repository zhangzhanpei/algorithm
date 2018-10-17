/**
 * 给定一个int数组，将数组中的偶数元素移动到数组前半部分
 */
class Solution {
    public int[] sortArrayByParity(int[] A) {
        int[] ret = new int[A.length];
        int i = 0, j = A.length - 1;
        for (int k : A) {
            if (k % 2 == 0) {
                ret[i] = k;
                i++;
            } else {
                ret[j] = k;
                j--;
            }
        }
        return ret;
    }
}
