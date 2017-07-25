//给定一个n个唯一元素的数组, 找到缺少的数. 如[0, 1, 3]时n=3, 则数组中缺少2
var missingNumber = function(nums) {
    let len = nums.length;
    let sum = 0;
    nums.map(function(n) {
        sum += n;
    });
    return len * (len + 1) / 2 - sum;
};

console.log(missingNumber([0, 1, 3])); //output 2
