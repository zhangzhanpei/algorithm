/**
 * 判断一个二叉树是否对称
 */
public boolean isSymmetric(TreeNode root) {
    return root == null || symmetric(root.left, root.right);
}

public boolean symmetric(TreeNode left, TreeNode right) {
    if (left == null || right == null) {
        return left == right;
    }
    if (left.val != right.val) {
        return false;
    }
    return symmetric(left.left, right.right) && symmetric(right.left, left.right);
}
