/**
 * 给定一个有序int数组，构建一个二叉查找树
 */
public TreeNode sortedArrayToBST(int[] nums) {
    if (nums.length == 0) return null;
    return getTreeNode(nums, 0, nums.length - 1);
}

private TreeNode getTreeNode(int[] nums, int start, int end) {
    if (start > end) return null;
    int mid = (int) Math.ceil((start + end) / 2.0); //取中点
    TreeNode node = new TreeNode(nums[mid]);
    node.left = getTreeNode(nums, start, mid - 1); //递归生成左子树和右子树
    node.right = getTreeNode(nums, mid + 1, end);
    return node;
}
