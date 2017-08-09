//给定一个数组, 将它两两分组, 取小的那个元素求和使得和最大
var arrayPairSum = function(nums) {
    nums.sort(function (a, b) { //先排序
        return a - b;
    });
    let i = 0, sum = 0;
    while (i < nums.length) {
        if (i % 2 === 0) {
            sum += nums[i];
        }
        i++;
    }
    return sum;
};
