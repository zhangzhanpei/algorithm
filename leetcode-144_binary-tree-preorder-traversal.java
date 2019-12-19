// 前序遍历二叉树(递归)
private List<Integer> ret = new ArrayList<>();
public List<Integer> preorderTraversal(TreeNode root) {
    if (root == null) {
        return this.ret;
    }
    this.ret.add(root.val);
    preorderTraversal(root.left);
    preorderTraversal(root.right);
    return this.ret;
}

// 前序遍历二叉树(栈)
public List<Integer> preorderTraversal(TreeNode root) {
    List<Integer> ret = new ArrayList<>();
    Stack<TreeNode> stack = new Stack<>();
    while (root != null || !stack.isEmpty()) {
        while (root != null) {
            ret.add(root.val); // 根节点先输出
            stack.push(root);
            root = root.left; // 左儿子成为根节点
        }

        if (!stack.isEmpty()) {
            root = stack.pop(); // 栈顶元素[最左下角的元素]出栈
            root = root.right; // 看有没有右孩子, 如果有则成为根节点继续输出和进栈
        }
    }
    return ret;
}
