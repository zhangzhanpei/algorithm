//Sqrt(x)
//给定一个整数x, 求x的平方根
//牛顿迭代法: 先猜测平方根为y. 如果y不对或精度不够, 令y=(y+x/y)/2, 可无限逼近正确值

function sqrt(x)
{
    var e = 1e-5; //精度准确到小数点后5位
    var y = 1;
    while (Math.abs(x - y * y) > e) {
        y = (y + x / y) / 2;
    }
    return Math.floor(y);
}
