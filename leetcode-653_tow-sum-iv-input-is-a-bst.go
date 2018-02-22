/**
 * 给定一个二叉树, 判断是否有两个节点的和等于给定值k
 * 前序遍历二叉树, 经过的节点写入map中, 遍历下一个节点时求得差值看在不在map中
 */
func findTarget(root *TreeNode, k int) bool {
    m := map[int]bool{}
    return dfs(root, m, k)
}

func dfs(root *TreeNode, m map[int]bool, k int) bool {
    if root == nil {
        return false
    }
    sub := k - root.Val //目标值减去当前节点得到差值
    if m[sub] { //判断差值是否在map中
        return true
    }
    m[root.Val] = true //如果不在, 把当前值加入到map
    return dfs(root.Left, m, k) || dfs(root.Right, m, k)
}
