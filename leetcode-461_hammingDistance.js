//Hamming Distance
//汉明距离, 两个等长字符串间对应位置的不同字符的个数
//先用异或得到不同字符的数量(1的个数), 然后统计1的个数即可
function hammingDistance(x, y)
{
    //异或运算, 0异或1得到1, 也就是结果中为1的地方, 就是两数分别为01(不同字符)的地方
    var xor = x ^ y;
    var count = 0;
    while (xor !== 0) {
        count++;
        //与运算, 每次运算都会消掉xor中最后的一个1
        xor = xor & (xor - 1);
    }
    return count;
}
