/**
 * 后序遍历二叉树
 */
var arr []int
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    postorderTraversal(root.Left)
    postorderTraversal(root.Right)
    arr = append(arr, root.Val)
    return arr
}
