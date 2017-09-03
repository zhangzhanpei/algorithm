//找出单词 t 多加的字符, t 是 s 随机打乱后加了一个字符形成的单词
//将单词逐个字符转成ascii值求和, 然后 t - s 即得到 t 多加的字符
var findTheDifference = function(s, t) {
    let charCodeS = 0, charCodeT = 0;
    for (let i = 0; i < s.length; i++) {
        charCodeS += s[i].charCodeAt();
    }
    for (let i = 0; i < t.length; i++) {
        charCodeT += t[i].charCodeAt();
    }
    return String.fromCharCode(charCodeT - charCodeS);
};

console.log(findTheDifference('abc', 'abcd'));
