/**
 * 给定一个二叉搜索树和一个范围[L, R], 移除二叉搜索树上不在范围内的元素, 返回新的二叉搜索树
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val < L {
        return trimBST(root.Right, L, R)
    }
    if root.Val > R {
        return trimBST(root.Left, L, R)
    }
    root.Left = trimBST(root.Left, L, R)
    root.Right = trimBST(root.Right, L, R)
    return root
}
