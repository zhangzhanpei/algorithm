//给定一个字符串, 找到第一个不重复字符的位置
var firstUniqChar = function(s) {
    let charArr = s.split(''); //先切割成字符数组
    for (let i = 0; i < charArr.length; i++) {
        if (i != 0) { //由前向后, 逐个字符交换到第一个元素位置
            let tmp = charArr[0];
            charArr[0] = charArr[i];
            charArr[i] = tmp;
        }
        if (charArr.lastIndexOf(charArr[0]) == 0) { //如果该字符的最后一次出现就是数组首部, 说明该字符唯一
            return i;
        }
    }
    return -1;
};
