/**
 * 给定一个二叉树和一个整数sum, 判断二叉树中是否有一条路径的元素的和为sum
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil { //到达叶子节点, 如果叶子节点的值等于sum剩下的值, 那么这条路径正确
        return sum == root.Val
    }

    //sum减去经过的节点的值, 直到叶子节点
    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}
