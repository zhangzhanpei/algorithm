/**
 * 中序遍历二叉树
 */
var arr []int
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    inorderTraversal(root.Left)
    arr = append(arr, root.Val)
    inorderTraversal(root.Right)
    return arr
}
