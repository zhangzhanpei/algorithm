/*
 * 自底向上逐层返回二叉树的元素
 */
public List<List<Integer>> levelOrderBottom(TreeNode root) {
    List<List<Integer>> ret = new ArrayList<>();
    List<TreeNode> l = new ArrayList<>(); //保存二叉树每层节点
    if (root == null) return ret;
    l.add(root);
    while (!l.isEmpty()) {
        List<Integer> tmp = new ArrayList<>();
        int size = l.size();
        for (int i = 0; i < size; i++) { //遍历一层节点取得值
            TreeNode node = l.get(0);
            tmp.add(node.val);
            if (node.left != null) l.add(node.left);
            if (node.right != null) l.add(node.right);
            l.remove(node); //遍历过的节点移除
        }
        ret.add(0, tmp);
    }
    return ret;
}
