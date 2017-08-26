/**
 * 给定一个二叉树, 找出根节点到叶子节点的所有路径
 */
var paths []string
func binaryTreePaths(root *TreeNode) []string {
    paths = []string{}
    if root != nil {
        searchPath(root, "")
    }
    return paths
}

func searchPath(node *TreeNode, path string) {
    if node.Left == nil && node.Right == nil { //到达叶子节点, 记录一条路径
        path += strconv.Itoa(node.Val)
        paths = append(paths, path)
    }
    if node.Left != nil { //继续处理左儿子
        searchPath(node.Left, path + strconv.Itoa(node.Val) + "->")
    }
    if node.Right != nil { //继续处理右儿子
        searchPath(node.Right, path + strconv.Itoa(node.Val) + "->")
    }
}
