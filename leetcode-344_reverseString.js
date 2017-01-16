//翻转字符串
function reverseString(s)
{
    var len = s.length;
    var rs = '';
    for (let i = len - 1; i >= 0; i--) {
        rs += s[i];
    }
    return rs;
}
