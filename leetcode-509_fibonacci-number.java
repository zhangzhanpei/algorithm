/**
 * 斐波那契数列，计算第N个数
 */
class Solution {
    public static void main(String[] args) {
        System.out.println(fib(4));
    }

    public static int fib(int N) {
        if (N == 0) {
            return 0;
        }
        if (N == 1) {
            return 1;
        }

        int a = 0;
        int b = 1;
        int sum = 0;

        for (int i = 2; i <= N; i++) {
            sum = a + b;
            a = b;
            b = sum;
        }

        return sum;
    }
}
