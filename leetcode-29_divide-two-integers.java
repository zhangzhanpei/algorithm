/**
 * 给定两个整数 dividend 和 divisor，不使用 乘法/除法/mod运算 求商
 * 除法可用减法实现，商就是减的次数
 */
public int divide(int dividend, int divisor) {
    if (dividend == Integer.MIN_VALUE && divisor == -1) {
        return Integer.MAX_VALUE;
    }
    if (dividend == 0) {
        return 0;
    }
    if (divisor == 1) {
        return dividend;
    }
    // 判断结果的正负号
    int sign = 1;
    if ((dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)) {
        sign = -1;
    }
    // 转成 long 防止溢出
    long dividendLong = Math.abs((long) dividend);
    long divisorLong = Math.abs((long) divisor);

    // 这里使用线性减法，当 dividend 和 divisor 相差很大时耗时较久
    // int ret = 0;
    // while (dividendLong >= divisorLong) {
    //     dividendLong -= divisorLong;
    //     ret++;
    // }

    // 这里使用指数减法，即 divisor 是每减去一次则翻倍
    int ret = 0;
    int multi = 1;
    long initDivisorLong = divisorLong;
    while (dividendLong >= divisorLong) {
        dividendLong -= divisorLong;
        ret += multi;

        divisorLong += divisorLong; // divisor 翻倍
        multi += multi; // 则减去的次数也应双倍计算

        if (dividendLong < divisorLong) { // 当 divisor 增大到大于剩下的 dividend 时，恢复到起初的 divisor 继续减
            divisorLong = initDivisorLong;
            multi = 1;
        }
    }

    if (sign < 0) {
        ret = -ret;
    }
    return ret;
}
