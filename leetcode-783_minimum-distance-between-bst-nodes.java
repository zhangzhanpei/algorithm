/**
 * 给定一个二叉搜索树，求任意节点间的最小差值
 */
public class Solution {
    private int min = Integer.MAX_VALUE;
    private Integer preVal = null; //递归时记录前一个节点值

    public int minDiffInBST(TreeNode root) {
        //中序遍历二叉树，逐个节点减去前一个节点，最小差值保存到min属性中
        if (root.left != null) minDiffInBST(root.left);
        if (this.preVal != null) this.min = Math.min(this.min, root.val - this.preVal);
        this.preVal = root.val;
        if (root.right != null) minDiffInBST(root.right);
        return this.min;
    }
}
