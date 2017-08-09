//计算小于n的素数个数[素数是只能被1和自身整除的数]
//2是最小的素数, 那么2的倍数均不为素数, 所以我们可以将小于n的数中2的倍数, 全部排除掉
//排除掉2的整数倍后, 剩下的数中大于2的最小的数就是下一个素数, 也就是3
//同样我们可以排除掉小于n的数中3的整数倍的数, 得到下一个素数为5, 如此循环
//因此我们声明一个n元素的数组, 默认没有赋值, 将素数的倍数位置设为0, 最后统计没赋值的元素个数即可

function countPrimes(n)
{
    if (n <= 2) {
        return 0;
    }
    var arr = new Array(n); //声明一个n元素的数组, 元素默认值为undefined
    var i = 2;
    while (i * i < n) { //里层的j是从i开始的, 也就是里层最小为i*i, 所以这里i*i要小于n, 否则就没必要再进到里层计算i的倍数了
        if (arr[i] === undefined) { //等于undefined说明i是素数
            var j = i;
            while (i * j < n) {
                arr[i * j] = 0; //把素数i的倍数位置的元素设为0
                j++;
            }
        }
        i++;
    }
    var count = 0;
    for (var x = 2; x < n; x++) {
        if (arr[x] === undefined) { //计算数组中undefined的个数即为素数的个数
            count++;
        }
    }
    return count;
}
