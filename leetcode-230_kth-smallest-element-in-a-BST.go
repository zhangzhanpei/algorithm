/**
 * 给定一个二叉查找树, 找到第k小的节点
 * 中序遍历二叉树, 遍历结果中的第k个元素即第k小的节点
 */
var arr []int
func kthSmallest(root *TreeNode, k int) int {
    arr = []int{}
    inorderTraversal(root)
    return arr[k - 1]
}

func inorderTraversal(node *TreeNode) {
    if node == nil {
        return
    }
    inorderTraversal(node.Left)
    arr = append(arr, node.Val)
    inorderTraversal(node.Right)
}
