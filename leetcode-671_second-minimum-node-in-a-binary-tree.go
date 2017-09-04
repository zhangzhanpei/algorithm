/**
 * 给定一个二叉树, 二叉树的每一个节点都有 0 个或 2 个子节点. 如果一个节点右两个子节点, 那么这个节点比两个子节点小
 * 找到整个二叉树中第二小的节点
 * 很明显根节点是最小的, 因此只需遍历二叉树, 找到除根节点外最小的节点即可
 */
var ret int
func findSecondMinimumValue(root *TreeNode) int {
    ret = math.MaxInt32
    traverse(root, root.Val)
    if ret == math.MaxInt32 { //如果这里相等说明其他节点都没有比 root 大
        return -1
    } else {
        return ret
    }
}

func traverse(node *TreeNode, r int) {
    if node == nil {
        return
    }
    //找到比 root 大但又是其他节点中最小的
    if r < node.Val && node.Val < ret {
        ret = node.Val
    }
    traverse(node.Left, r)
    traverse(node.Right, r)
}
