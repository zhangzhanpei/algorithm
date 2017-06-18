//给定一个字符串, 找到最长不重复字串的长度
function lengthOfLongestSubstring(s) {
    let tmpArr = []; //不重复字符存入数组中
    let max = 0;
    let start = 0; //记录开始位置
    let len = s.length;
    for (let i = 0; i < len; i++) {
        if (tmpArr.indexOf(s[i]) > -1) { //如果字符已经在数组中, 出现重复
            if (tmpArr.length > max) {
                max = tmpArr.length; //设置新的最大长度
            }
            i = ++start - 1; //开始位置+1, 重置i
            tmpArr = []; //重置数组
        } else {
            tmpArr.push(s[i]);
        }
    }
    return tmpArr.length > max ? tmpArr.length : max;
};

let str = "abcabcbb";
console.log(lengthOfLongestSubstring(str));
