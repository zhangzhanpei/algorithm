//给定整数n, 翻转它的二进制表示
//n%2取余数, 越早的余数, 就是翻转后的越高位. 如果余数为1, 乘上它的加权, 然后求和

function reverseBits(n)
{
    var rn = 0;
    for (var i = 32; i > 0; i--) {
        var j = n % 2;
        n = parseInt(n / 2);
        if (j == 1) {
            var tmp = 1;
            for (var x = 1; x < i; x++) {
                tmp = tmp * 2;
            }
            rn += tmp;
        }
    }
    return rn;
}
