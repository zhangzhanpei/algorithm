/**
 * 给定一个二叉树, 求相同节点值的路径最大长度
 */
var ret int

func longestUnivaluePath(root *TreeNode) int {
    ret = 0
    recursive(root)
    return ret
}

func recursive(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := recursive(root.Left)
    right := recursive(root.Right)
    leftPath, rightPath := 0, 0
    if root.Left != nil && root.Val == root.Left.Val {
        leftPath = left + 1
    }
    if root.Right != nil && root.Val == root.Right.Val {
        rightPath = right + 1
    }
    ret = int(math.Max(float64(ret), float64(leftPath + rightPath)))
    return int(math.Max(float64(leftPath), float64(rightPath)))
}
