/**
 * 给定一个二叉树，求每层的最大值
 */
public List<Integer> largestValues(TreeNode root) {
    List<Integer> ret = new ArrayList<>();
    if (root == null) {
        return ret;
    }
    ArrayList<TreeNode> list = new ArrayList<>();
    list.add(root);
    while (list.size() > 0) {
        int len = list.size();
        int max = -2147483648;
        for (int i = 0; i < len; i++) { //广度优先算法
            TreeNode node = list.remove(0);
            if (node.val > max) { //取最大值
                max = node.val;
            }
            if (node.left != null) {
                list.add(node.left);
            }
            if (node.right != null) {
                list.add(node.right);
            }
        }
        ret.add(max);
    }
    return ret;
}
