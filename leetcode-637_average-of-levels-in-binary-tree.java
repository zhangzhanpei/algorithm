/**
 * 给定一个二叉树，求每一层的平均值
 */
public List<Double> averageOfLevels(TreeNode root) {
    List<Double> ret = new ArrayList<>();
    if (root == null) {
        return ret;
    }
    ArrayList<TreeNode> list = new ArrayList<>();
    list.add(root);
    while (list.size() > 0) {
        int len = list.size();
        double sum = 0;
        for (int i = 0; i < len; i++) { //使用广度优先算法
            TreeNode node = list.remove(0);
            sum += node.val;
            if (node.left != null) {
                list.add(node.left);
            }
            if (node.right != null) {
                list.add(node.right);
            }
        }
        ret.add(sum / len); //求均值
    }
    return ret;
}
