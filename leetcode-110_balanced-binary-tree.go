/**
 * 给定一个二叉树, 判断是否平衡二叉树
 * 平衡二叉树：空树, 或左右子树的高度差不超过1, 并且左右子树都是平衡二叉树
 */
func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return int(math.Max(float64(depth(root.Left)), float64(depth(root.Right)))) + 1
}

func isBalanced(root *TreeNode) bool {
    //空树
    if root == nil {
        return true
    }
    leftDepth := depth(root.Left)
    rightDepth := depth(root.Right)
    //左右子树的高度差不超过1, 并且左右子树都是平衡二叉树
    return math.Abs(float64(leftDepth - rightDepth)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
