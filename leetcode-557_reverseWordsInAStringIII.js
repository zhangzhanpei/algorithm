//给定一个字符串, 逐个单词反转
var reverseWords = function(s) {
    let arr = s.split(" "); //先切割成单词数组
    let i = 0;
    while (arr[i]) {
        arr[i] = arr[i].split("").reverse().join(""); //逐个单词反转
        i++;
    }
    return arr.join(" "); //再拼接回来
};

reverseWords("Let's take LeetCode contest"); // output: "s'teL ekat edoCteeL tsetnoc"
