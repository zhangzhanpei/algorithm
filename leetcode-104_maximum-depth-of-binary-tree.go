/**
 * 给定一个二叉树, 求二叉树的最大高度
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftMaxDepth := maxDepth(root.Left) + 1
    rightMaxDepth := maxDepth(root.Right) + 1
    return int(math.Max(float64(leftMaxDepth), float64(rightMaxDepth)))
}
