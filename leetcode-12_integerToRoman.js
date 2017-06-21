//int 转罗马数字
var intToRoman = function(num) {
    let m = {
        1: "I",
        2: "II",
        3: "III" ,
        4: "IV",
        5: "V",
        6: "VI",
        7: "VII",
        8: "VIII" ,
        9: "IX",
        10: "X",
        20: "XX",
        30: "XXX",
        40: "XL",
        50: "L",
        60: "LX",
        70: "LXX",
        80: "LXXX",
        90: "XC",
        100: "C",
        200: "CC",
        300: "CCC",
        400: "CD",
        500: "D",
        600: "DC",
        700: "DCC",
        800: "DCCC",
        900: "CM",
        1000: "M",
        2000: "MM",
        3000: "MMM"
    };
    let diviser = 1000; //除数设为1000, 假设 num 最大为 3999
    let res = ""; //返回的罗马数字
    let i = 0;
    while (num > 0) {
        i = parseInt(num / diviser); //取整
        if (m[i * diviser] !== undefined) { //去 m 中去相应的罗马符号
            res += m[i * diviser];
        }
        num %= diviser; //num 设为余数
        diviser /= 10; //除数相应除以10, 继续处理余数部分
    }
    return res;
};

console.log(intToRoman(1));
