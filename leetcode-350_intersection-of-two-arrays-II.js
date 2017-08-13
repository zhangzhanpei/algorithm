//求两个数组的交集[结果元素可重复]
var intersect = function(nums1, nums2) {
    //统计nums1中各个元素出现的次数
    let map = new Map();
    nums1.map(function(n) {
        if (map[n] !== undefined) {
            map[n] += 1;
        } else {
            map[n] = 1;
        }
    });

    //从nums2中取相应次数的元素即是交集
    return nums2.filter(function(n) {
        if (map[n]) {
            map[n]--;
            return true;
        } else {
            return false;
        }
    });
};

intersect([1, 2, 2, 1], [2, 2]);
