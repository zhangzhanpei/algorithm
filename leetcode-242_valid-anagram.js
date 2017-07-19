//给定两个单词s和t, 判断t是否为s的乱序
//使用hashtable, 从s读到的字符+1, 从t读到的字符-1, 最后应该所有字符的数量刚好抵消
var isAnagram = function(s, t) {
    if (s.length !== t.length) return false;
    let m = new Map();
    let n = s.length;
    for (let i = 0; i < n; i++) {
        if (m[s[i]] !== undefined) {
            m[s[i]] += 1;
        } else {
            m[s[i]] = 1;
        }
        if (m[t[i]] !== undefined) {
            m[t[i]] -= 1;
        } else {
            m[t[i]] = -1;
        }
    }
    console.log(m);
    for (let key in m) {
        if (m[key] !== 0) {
            return false;
        }
    }
    return true;
};

console.log(isAnagram("a", "b"));
