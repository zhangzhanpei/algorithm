//给定一个单词数组, 返回字母都在键盘同一行的单词
var findWords = function(words) {
    let row1 = 'QWERTYUIOP';
    let row2 = 'ASDFGHJKL';
    let row3 = 'ZXCVBNM';
    let i = 0;
    let res = [];
    while (words[i]) {
        let word = words[i].toUpperCase();
        let len = word.length;
        let j = 0;
        let row; //先判断单词首字母位于键盘哪一行
        if (row1.indexOf(word[0]) >= 0) {
            row = row1;
        }
        if (row2.indexOf(word[0]) >= 0) {
            row = row2;
        }
        if (row3.indexOf(word[0]) >= 0) {
            row = row3;
        }
        while (word[j]) {
            if (row.indexOf(word[j]) == -1) { //确定首字母的行后, 遍历后面的字母看是否也在同一行
                break;
            }
            j++;
            if (j == len) { //如果全在同一行
                res.push(words[i]);
            }
        }
        i++; //下一个单词
    }
    return res;
};

findWords(["Hello", "Alaska", "Dad", "Peace"]); //output: ["Alaska", "Dad"]
