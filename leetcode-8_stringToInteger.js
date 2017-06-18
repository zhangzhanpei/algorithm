//字符串转数字
var myAtoi = function(str) {
    let num = 0;
    let flag = 1; //正负数标志
    let max = 2147483647;
    let min = -2147483648;
    str = str.trim();
    if (['+', '-'].indexOf(str[0]) > -1) { //如果开头是+-号, 从第二位开始处理
        if (str[0] === '-') {
            flag = -1;
        }
        str = str.substr(1);
    }

    let len = str.length;
    for (let i = 0; i < len; i++) {
        let tmp = parseInt(str[i]);
        if (isNaN(tmp)) { //如果不是数字, 看情况返回
            if (num === 0) { //如果值是0直接返回
                return num;
            } else { //如果值不是0则跳出循环去处理一下正负
                break;
            }
        }
        num *= 10;
        num += tmp;
    }
    num *= flag;
    if (num > max) {
        return max;
    } else if (num < min) {
        return min;
    }
    return num;
};

let a = "  -0012a42";
console.log(myAtoi(a));
