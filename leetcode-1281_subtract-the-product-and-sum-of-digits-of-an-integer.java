/**
 * 给定一个整数 n，则各位数乘积为 i，各位数总和为 j，求 i - j
 * 如给定 n = 234，则 i = 2 * 3 * 4 = 24，j = 2 + 3 + 4 = 9，则结果为 24 - 9 = 15
 */
class Solution {
    public static void main(String[] args) {
        int n = 234;
        System.out.println(subtractProductAndSum(n));
    }
    public static int subtractProductAndSum(int n) {
        int product = 1;
        int sum = 0;
        while (n > 0) {
            int remainder = n % 10;
            n = n / 10;
            product *= remainder; // 各位数乘积
            sum += remainder; // 各位数总和
        }
        return product - sum;
    }
}
