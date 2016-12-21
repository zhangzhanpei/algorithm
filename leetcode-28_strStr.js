//查找字符串是否包含子串, 返回第一个子串的起始位置或-1
//从子串的第一个字符开始比较, 如果某个字符不同, 则开始下一轮
//如果读取到子串尾部的下一位, 此时越界会得到undefined, 说明子串完全匹配, 返回haystack字符串的位置i即可
function strStr(haystack, needle)
{
    for (var i = 0; i <= haystack.length; i++) {
        for (var j = 0; j <= needle.length; j++) {
            if (needle[j] === undefined) {
                return i;
            }
            if (haystack[i + j] === undefined) {
                return -1;
            }
            if (haystack[i + j] != needle[j]) {
                break;
            }
        }
    }
}

var haystack = 'Hello, World!';
var needle = 'o';
console.log(strStr(haystack, needle));
