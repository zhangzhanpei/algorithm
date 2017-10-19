/**
 * 给定一个二叉树和一个整数sum, 判断二叉树中元素的和为sum的路径数(路径不用抵达叶子节点)
 */
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    //根节点路径数 + 左右儿子节点的路径数
    return hasPathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func hasPathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    if sum == root.Val { //找到一条路径
        //这里继续求左右儿子节点路径数是因为节点是正负数的, 左右儿子节点可能拥有和为0的路径
        return 1 + hasPathSum(root.Left, sum - root.Val) + hasPathSum(root.Right, sum - root.Val)
    } else {
        return hasPathSum(root.Left, sum - root.Val) + hasPathSum(root.Right, sum - root.Val)
    }
}
