/**
 * 给定一个二叉树, 求二叉树的最大高度
 */
public int maxDepth(TreeNode root) {
    if (root == null) {
        return 0;
    }

    int leftDepth = maxDepth(root.left) + 1;
    int rightDepth = maxDepth(root.right) + 1;

    return Math.max(leftDepth, rightDepth);
}
