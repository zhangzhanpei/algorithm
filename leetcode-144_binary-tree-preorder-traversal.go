/**
 * 前序遍历二叉树
 */
var arr []int
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    arr = append(arr, root.Val)
    preorderTraversal(root.Left)
    preorderTraversal(root.Right)
    return arr
}
