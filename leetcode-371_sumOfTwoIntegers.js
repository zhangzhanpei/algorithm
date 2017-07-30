//不使用+-号求两个整数的和
var getSum = function(a, b) {
    if (b === 0) return a;
    let sum = 0, carry = 0;
    sum = a ^ b;
    carry = (a & b) << 1;
    return getSum(sum, carry);
};

console.log(getSum(2, 3));
