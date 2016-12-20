//在一个有序数组中找出两个元素之和等于给定值
function twoSum(nums, target)
{
    var i = 0;
    var j = nums.length - 1;
    while (i < j) {
        let sum = nums[i] + nums[j];
        if (sum == target) {
            return [i, j];
        } else if (sum < target) {
            i++;
        } else {
            j--;
        }
    }
    return [-1, -1];
}
