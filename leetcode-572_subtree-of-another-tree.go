/**
 * 给定两个二叉树 s 和 t, 判断 t 是否为 s 的子树
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {
        return false
    }
    if isSame(s, t) {
        return true
    }
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil || t == nil {
        return false
    }
    if s.Val != t.Val {
        return false
    }
    return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
