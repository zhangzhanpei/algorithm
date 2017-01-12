//给出一个正整数n, 计算n!后面0的个数
//偶数*5会在后面产生0, 而偶数远比5的个数多, 所以只要计算5的分解的个数就好了

function trailingZeroes(n)
{
    var count = 0;
    while (n) {
        count += Math.floor(n/5);
        n = Math.floor(n/5);
    }
    return count;
}
