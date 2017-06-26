//合并两个二叉树[相同位置节点值相加]
var mergeTrees = function(t1, t2) {
    //当两棵树相同位置都没有节点时, return null
    if (t1 || t2) {
        let node = TreeNode((t1 && t1.val || 0) + (t2 && t2.val || 0)); //相同位置节点值相加
        node.left = mergeTrees(t1 && t1.left, t2 && t2.left);
        node.right = mergeTrees(t1 && t1.right, t2 && t2.right);
        return node;
    } else {
        return null;
    }
};
