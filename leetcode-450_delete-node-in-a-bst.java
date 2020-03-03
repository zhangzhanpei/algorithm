/**
 * 删除二叉查找树的结点
 */
class Solution {
    public TreeNode deleteNode(TreeNode root, int key) {
        if (root == null) {
            return null;
        }
        if (key < root.val) { // 如果要删除的值比当前结点小，递归去左子树删除结点
            root.left = deleteNode(root.left, key);
        } else if (key > root.val) { // 如果要删除的值比当前结点大，递归去右子树删除结点
            root.right = deleteNode(root.right, key);
        } else { // 如果刚好是要删除当前结点
            // 只有左子树或只有右子树或左右子树都没有
            if (root.left == null || root.right == null) {
                root = (root.left != null) ? root.left : root.right;
            } else { // 有左右子树
                TreeNode curr = root.right;
                while (curr.left != null) { // 在右子树中找到后继结点（右子树中的最小值，即右子树最左下角那个结点）
                    curr = curr.left;
                }
                root.val = curr.val; // 当前要删除的结点赋值为后继结点的值
                root.right = deleteNode(root.right, curr.val); // 再去右子树删除后继结点（递归）
            }
        }
        return root;
    }
}
