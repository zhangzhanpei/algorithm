/**
 * 给定两个二叉查找树，取出两棵树所有结点按升序返回
 */
public List<Integer> getAllElements(TreeNode root1, TreeNode root2) {
    List<Integer> ret = new LinkedList<>();
    Stack<TreeNode> stack1 = new Stack<>(), stack2 = new Stack<>();
    putLeft(stack1, root1);
    putLeft(stack2, root2);
    while (!stack1.isEmpty() && !stack2.isEmpty()) {
        TreeNode node1 = stack1.peek(), node2 = stack2.peek();
        if (node1.val <= node2.val) {
            node1 = stack1.pop();
            ret.add(node1.val);
            putLeft(stack1, node1.right); // 右儿子进行前序遍历
        } else {
            node2 = stack2.pop();
            ret.add(node2.val);
            putLeft(stack2, node2.right);
        }
    }
    
    // 继续处理剩下的栈元素
    while (!stack1.isEmpty()) {
        TreeNode node = stack1.pop();
        ret.add(node.val);
        putLeft(stack1, node.right);
    }
    while (!stack2.isEmpty()) {
        TreeNode node = stack2.pop();
        ret.add(node.val);
        putLeft(stack2, node.right);
    }

    return ret;
}

// 前序遍历，栈顶即最小的结点
public void putLeft(Stack<TreeNode> stack, TreeNode root) {
    while (root != null) {
        stack.push(root);
        root = root.left;
    }
}
