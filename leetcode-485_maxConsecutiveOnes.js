//给定一个0和1组成的数组, 找出连续的1的最大个数
var findMaxConsecutiveOnes = function(nums) {
    let max = 0, tmp = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === 0) {
            max = Math.max(max, tmp);
            tmp = 0;
        } else {
            tmp += 1;
        }
    }
    max = Math.max(max, tmp);
    return max;
};

console.log(findMaxConsecutiveOnes([1,1,0,1,1,1])); //output 3
