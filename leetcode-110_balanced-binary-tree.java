/**
 * 判断二叉树是否平衡二叉树
 */
class Solution {
    private boolean balance = true;

    // 深度优先遍历，自底向上判断子树是否平衡，避免自顶向下递归的重复计算高度
    public int depth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leftDepth = depth(root.left);
        int rightDepth = depth(root.right);
        if (Math.abs((leftDepth - rightDepth)) > 1) {
            balance = false;
        }
        return Math.max(leftDepth, rightDepth) + 1;
    }
    public boolean isBalanced(TreeNode root) {
        depth(root);
        return balance;
    }
}
