/**
 * 给定一个二叉查找树, 转成一棵更大的树
 */
class Solution {
    private int sum = 0;

    public TreeNode convertBST(TreeNode root) {
        if (root == null) {
            return null;
        }
        convertBST(root.right);
        root.val += this.sum;
        this.sum = root.val;
        convertBST(root.left);
        return root;
    }
}
