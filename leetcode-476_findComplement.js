//给定一整数n, 将n转成2进制, 去除前面的0, 然后按位取反, 再转回10进制

function findComplement(num)
{
    var tmp = num.toString(2);
    var len = tmp.length;
    var rn = 0;
    for (var i = len - 1; i >= 0; i--) {
        if (tmp[i] == '0') {
            rn += Math.pow(2, len - i - 1);
        }
    }
    return rn;
}
