/**
 * 按层遍历二叉树, 每层节点以数组形式返回
 * BFS广度优先遍历
 */
public List<List<Integer>> levelOrder(TreeNode root) {
    List<List<Integer>> ret = new ArrayList<>();
    if (root == null) {
        return ret;
    }
    // 存储每层的结点
    List<TreeNode> list = new LinkedList<>();
    list.add(root);
    while (list.size() > 0) {
        int len = list.size();
        // 存储每层的结点的值
        List<Integer> tmp = new ArrayList<>();
        for (int i = 0; i < len; i++) {
            // 弹出头元素
            TreeNode treeNode = list.remove(0);
            tmp.add(treeNode.val);
            if (treeNode.left != null) {
                list.add(treeNode.left);
            }
            if (treeNode.right != null) {
                list.add(treeNode.right);
            }
        }
        // 遍历完一层
        ret.add(tmp);
    }
    return ret;
}
