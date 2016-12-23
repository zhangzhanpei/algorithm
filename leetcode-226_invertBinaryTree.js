//翻转二叉树
//递归交换左右节点即可
function invertBinaryTree(root)
{
    if (root === null) {
        return null;
    }
    tmp = root.left;
    root.left = root.right;
    root.right = tmp;
    invertBinaryTree(root.left);
    invertBinaryTree(root.right);
    return root;
}
