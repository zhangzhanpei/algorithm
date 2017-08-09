//罗马数字转int
var romanToInt = function(s) {
    let map = {I: 1, V: 5, X: 10, L: 50, C: 100, D: 500, M: 1000};
    let num = 0;
    let len = s.length;
    for (let i = 0; i < len - 1; i++) {
        let c = s[i];
        let next = s[i + 1];
        if (map[c] < map[next]) { //当属性名为变量时使用中括号访问
            num -= map[c]; //罗马数字当左边比较小时做减法
        } else {
            num += map[c]; //当左边数字比较大时做加法
        }
    }

    num += map[s[len - 1]]; //加上最后一位
    return num;
};

let roman = "IV";
console.log(romanToInt(roman));
