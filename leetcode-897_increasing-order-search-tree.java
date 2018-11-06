/**
 * 给定一个二叉查找树，重构成一个只有右节点且仍然有序的二叉树
 */
class Solution {
    // 存储所有节点的值
    private List<Integer> nodes = new ArrayList<>();

    public TreeNode increasingBST(TreeNode root) {
        // 中序遍历，得到的数组是升序的
        this.inorderTraversal(root);
        // 使用数组构建二叉树
        TreeNode ret = new TreeNode(0), currentNode = ret;
        for (Integer i : this.nodes) {
            currentNode.right = new TreeNode(i);
            currentNode = currentNode.right;
        }
        return ret.right;
    }

    public void inorderTraversal(TreeNode root) {
        if (root == null) {
            return;
        }
        inorderTraversal(root.left);
        this.nodes.add(root.val);
        inorderTraversal(root.right);
    }
}
