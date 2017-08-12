//求两个数组的交集[结果元素不可重复]
var intersection = function(nums1, nums2) {
    let res = [];
    for (let i = 0; i < nums1.length; i++) {
        if (res.indexOf(nums1[i]) === -1 && nums2.indexOf(nums1[i]) >= 0) {
            res.push(nums1[i]);
        }
    }
    return res;
};
