/**
 * 给定一个二叉树，返回最左下角的节点的值
 */
public int findBottomLeftValue(TreeNode root) {
    ArrayList<TreeNode> list = new ArrayList<>();
    list.add(root);
    TreeNode node = null;
    while (list.size() > 0) {
        node = list.remove(0); //第一个元素出队
        if (node.right != null) { //右儿子先进队
            list.add(node.right);
        }
        if (node.left != null) { //左儿子后进队，那么最左下方的节点时最晚进队的
            list.add(node.left);
        }
    }
    return node.val;
}