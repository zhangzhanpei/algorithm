/**
 * 翻转二叉树
 * 递归交换左右节点即可
 */
public TreeNode invertTree(TreeNode root) {
    if (root == null) return null;
    //交换左右子节点
    TreeNode tmp = root.left;
    root.left = root.right;
    root.right = tmp;
    //递归处理左右子节点
    invertTree(root.left);
    invertTree(root.right);
    return root;
}
