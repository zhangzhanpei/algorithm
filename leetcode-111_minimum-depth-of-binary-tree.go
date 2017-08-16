/**
 * 给定一个二叉树, 求二叉树的最小高度
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1
}
