//leetcode-1
//无序数组中找出两个元素使得它们的和等于给定值
//使用一个map保存差值, 逐个读取数组元素, 如果元素值在map的键中, 说明该元素就是所需的差值
//此时可从map中读取到当初产生该差值的元素位置
function twoSum(nums, target) {
    var map = {};
    for (var k in nums) {
        if (map[nums[k]] === undefined) {
            map[target - nums[k]] = k;
        } else {
            return [parseInt(map[nums[k]]), parseInt(k)];
        }
    }
}
