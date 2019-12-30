/**
 * 给定一个 int 数组，返回数组元素长度为偶数的元素数量
 */
public int findNumbers(int[] nums) {
    int ret = 0;
    for (int n : nums) {
        if (Integer.toString(n).length() % 2 == 0) {
            ret++;
        }
    }
    return ret;
}
