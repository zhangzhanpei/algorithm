/**
 * 给定一个二叉树，判断是否为二叉搜索树
 */
class Solution {

    private TreeNode prev = null;

    // 方案一：递归
    public boolean isValidBST(TreeNode root) {
        if (root == null) {
            return true;
        }

        if (!isValidBST(root.left)) {
            return false;
        }

        // 先处理左子树，所以 prev 是左子树的结点
        if (prev != null && root.val <= prev.val) {
            return false;
        }

        prev = root;

        if (!isValidBST(root.right)) {
            return false;
        }

        return true;
    }

    // 方案二：栈
    public boolean isValidBST(TreeNode root) {
        if (root == null) {
            return true;
        }
        Stack<TreeNode> stack = new Stack<>();
        TreeNode prev = null;
        // 这里使用中序遍历二叉树
        while (root != null || !stack.isEmpty()) {
            while (root != null) {
                stack.push(root);
                root = root.left;
            }

            root = stack.pop(); // 第一个出栈的是左下角的叶子结点
            if (prev != null && root.val <= prev.val) {
                return false;
            }
            prev = root;
            root = root.right;
        }
        return true;
    }
}
