import java.util.Stack;

/**
 * 给定一个二叉搜索树, 实现 next 方法由小到大迭代返回二叉树节点
 */
class BSTIterator {

    private Stack<TreeNode> stack = new Stack<>();

    public BSTIterator(TreeNode root) {
        // 实例化对象时先沿着左儿子进栈
        while (root != null) {
            stack.push(root);
            root = root.left;
        }
    }

    /** @return the next smallest number */
    public int next() {
        // 此时栈顶元素为最小
        TreeNode root = stack.pop();
        int ret = root.val;
        // 如果有右儿子，同样沿着左儿子进栈
        if (root.right != null) {
            root = root.right;
            while (root != null) {
                stack.push(root);
                root = root.left;
            }
        }
        return ret;
    }

    /** @return whether we have a next smallest number */
    public boolean hasNext() {
        return stack.size() > 0;
    }
}
