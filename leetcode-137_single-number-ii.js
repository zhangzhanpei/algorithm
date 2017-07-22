//给定一个包含多个int的数组, 每个元素都出现3次, 但其中一个元素没有出现3次, 找到这个元素
var singleNumber = function(nums) {
    let m = new Map();
    for (let k in nums) {
        if (m[nums[k]] !== undefined) {
            m[nums[k]] += 1;
        } else {
            m[nums[k]] = 1;
        }
    }
    for (let i in m) {
        if (m[i] < 3) {
            return parseInt(i);
        }
    }
};
