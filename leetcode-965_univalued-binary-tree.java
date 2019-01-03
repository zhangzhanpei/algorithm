/**
 * 给定一个int二叉树，判断二叉树所有节点值是否相同
 */
class Solution {
    public int tmp;

    public boolean isUnivalTree(TreeNode root) {
        this.tmp = root.val;
        return this.preOrderTraversal(root);
    }

    // 前序遍历二叉树结点，与前值比较
    public boolean preOrderTraversal(TreeNode root) {
        if (root == null) {
            return true;
        }
        if (root.val != this.tmp) {
            return false;
        }
        return preOrderTraversal(root.left) && preOrderTraversal(root.right);
    }
}
