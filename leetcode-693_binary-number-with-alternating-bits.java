/**
 * 给定一个整数, 判断它的二进制表示是否由01交替显示
 * input 5  = 101,  return true
 * input 7  = 111,  return false
 * input 11 = 1011, return false
 * input 10 = 1010, return true
 */
public boolean hasAlternatingBits(int n) {
    // 先求前一次余数
    int prev = n % 2;
    n /= 2;
    while (n > 0) {
        // 再求一次余数
        int next = n % 2;
        // 如果两次余数一致则不是交替显示
        if (prev == next) {
            return false;
        }
        prev = next;
        n /= 2;
    }
    return true;
}
