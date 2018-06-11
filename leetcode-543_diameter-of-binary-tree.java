/**
 * 求二叉树最远的两个节点间的距离
 */
class Solution {
    private int max;

    public int diameterOfBinaryTree(TreeNode root) {
        maxDepth(root);
        return this.max;
    }

    /**
     * 对于任意节点作为根节点，两个节点间最远的距离即 左子树最大高度 + 右子树最大高度
     * 遍历二叉树节点计算即可
     */
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leftMaxDepth = maxDepth(root.left);
        int rightMaxDepth = maxDepth(root.right);
        // 替换最大距离
        this.max = Integer.max(leftMaxDepth + rightMaxDepth, this.max);
        // 返回当前节点作为根节点的最大高度
        return Integer.max(leftMaxDepth, rightMaxDepth) + 1;
    }
}
