//将有序数组转成二叉搜索树
var sortedArrayToBST = function(nums) {
    if (nums.length === 0) {
        return null;
    }
    let mid = parseInt(nums.length / 2);
    let node = new TreeNode(nums[mid]);
    node.left = sortedArrayToBST(nums.slice(0, mid));
    node.right = sortedArrayToBST(nums.slice(mid + 1));
    return node;
};
