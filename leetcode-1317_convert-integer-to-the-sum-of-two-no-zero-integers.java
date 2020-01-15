/**
 * 给定一个正整数 n，拆成两个整数 i 和 j 使得 i + j = n 且 i 和 j 各个位都不为 0 
 */
public int[] getNoZeroIntegers(int n) {
    for (int i = 1; i < n; i++) {
        if (!containZero(i) && !containZero(n - i)) {
            return new int[]{i, n - i};
        }
    }
    return new int[0];
}

public boolean containZero(int n) {
    while (n > 0) {
        if (n % 10 == 0) {
            return true;
        }
        n /= 10;
    }
    return false;
}
