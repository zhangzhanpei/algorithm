/**
 * 给定一个整数，判断是否回文的
 */
public boolean isPalindrome(int x) {
    // 负数或被10整除的数反过来是负号或0开头，不算回文
    if (x < 0 || x % 10 == 0) {
        return false;
    }
    int originX = x;
    int y = 0;
    while (x != 0) {
        y = y * 10 + x % 10;
        x /= 10;
    }

    return originX == y;
}
