//判断两个字符串是否结构相同
//给出"egg", "add", return true
//给出"foo", "bar", return false
//给出"paper", "title", return true

function isIsomorphic(s, t)
{
    if (s === null || t === null) {
        return false;
    }
    if (s.length != t.length) {
        return false;
    }
    var i = 0;
    var ms = {}; //两个字典, 保存字符的对应关系
    var mt = {};
    while (i < s.length) {
        if (ms[s[i]] === undefined && mt[t[i]] === undefined) {
            ms[s[i]] = t[i];
            mt[t[i]] = s[i];
        } else {
            if (ms[s[i]] != t[i] || mt[t[i]] != s[i]) {
                return false;
            }
        }
        i++;
    }
    return true;
}
