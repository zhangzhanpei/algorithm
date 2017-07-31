//找到数组中唯一有重复的数字
//遍历数组, 有重复数字的元素最后出现的位置必定不等于当前位置
var findDuplicate = function(nums) {
    for (let key in nums) {
        if (nums.lastIndexOf(nums[key]) != key) {
            return nums[key];
        }
    }
};

console.log(findDuplicate([1, 2, 3, 4, 4, 5]));
