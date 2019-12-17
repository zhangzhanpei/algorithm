/**
 * 给定一个整数 x, 求平方根
 * 牛顿迭代法: 先猜测平方根为 y. 如果 y 不对或精度不够, 令 y=(y+x/y)/2, 可无限逼近正确值
 */
class Solution {
    public static void main(String[] args) {
        System.out.println(mySqrt(8));;
    }

    public static int mySqrt(int x) {
        double y = 1.0;
        while (Math.abs(x - y * y) > 0.01) {
            y = (y + x / y) / 2;
        }
        return (int) y;
    }
}
