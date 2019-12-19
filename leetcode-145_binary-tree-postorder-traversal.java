// 后序遍历二叉树(递归)
private List<Integer> ret = new ArrayList<>();
public List<Integer> postorderTraversal(TreeNode root) {
    if (root == null) {
        return this.ret;
    }
    postorderTraversal(root.left);
    postorderTraversal(root.right);
    this.ret.add(root.val);
    return this.ret;
}

// 后序遍历二叉树(栈)
public List<Integer> postorderTraversal(TreeNode root) {
    List<Integer> ret = new ArrayList<>();
    if (root == null) {
        return ret;
    }
    TreeNode current = null, prev = null;
    Stack<TreeNode> stack = new Stack<>();
    stack.push(root);
    while (!stack.isEmpty()) {
        current = stack.peek();
        // 如果没有左右孩子, 或左右孩子都被访问过, 则输出当前节点值
        if ((current.left == null && current.right == null) || (prev != null && (prev == current.left || prev == current.right))) {
            ret.add(current.val);
            stack.pop();
            prev = current;
        } else {
            // 先把右孩子进栈, 再把左孩子进栈, 那么出栈顺序为左右中, 即后序遍历
            if (current.right != null) {
                stack.push(current.right);
            }
            if (current.left != null) {
                stack.push(current.left);
            }
        }
    }
    return ret;
}
