//给定一个整数 n, 从 1 到 n 如果是3的倍数就输出 Fizz, 是 5 的倍数就输出 Buzz, 同时被 3 和 5 整除则输出 FizzBuzz
var fizzBuzz = function(n) {
    let ret = [];
    for (let i = 1; i <= n; i++) {
        if (i % 15 === 0) {
            ret.push('FizzBuzz');
        } else if (i % 3 === 0) {
            ret.push('Fizz');
        } else if (i % 5 === 0) {
            ret.push('Buzz');
        } else {
            ret.push(`${i}`);
        }
    }
    return ret;
};

console.log(fizzBuzz(15));
