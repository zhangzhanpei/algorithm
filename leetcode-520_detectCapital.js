//判断一个单词的大小写是否正确
//全体大写, 全体小写, 首字母大写 3中均为正确
var detectCapitalUse = function(word) {
    if (word[0].toUpperCase() == word[0]) { //首字母大写
        if (word.slice(1).toUpperCase() == word.slice(1) || word.slice(1).toLowerCase() == word.slice(1)) {
            return true;
        } else {
            return false;
        }
    } else { //首字母小写
        if (word.toLowerCase() == word) {
            return true;
        } else {
            return false;
        }
    }
};

console.log(detectCapitalUse('Hello'));
