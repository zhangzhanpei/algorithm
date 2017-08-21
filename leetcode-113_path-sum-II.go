/**
 * 给定一个二叉树和一个整数sum, 找出二叉树中元素和为sum的所有路径
 */
var ret [][]int
func pathSum(root *TreeNode, sum int) [][]int {
    ret = [][]int{} //记录所有路径
    onePath := []int{} //记录一条路径
    onePathSum(root, sum, onePath)
    return ret
}

func onePathSum(root *TreeNode, sum int, onePath []int) {
    if root == nil {
        return
    }
    onePath = append(onePath, root.Val)
    if root.Left == nil && root.Right == nil && sum == root.Val { //到达叶子节点, 如果叶子节点的值等于sum剩下的值, 那么这条路径正确
        //这里需要复制一个新的onePath值push到ret里面, 否则继续后面的查找会改变onePath的值导致ret不正确(onePath传引用)
        tmp := make([]int, len(onePath))
        copy(tmp, onePath) 
        ret = append(ret, tmp)
        return
    }

    //sum减去经过的节点的值, 直到叶子节点
    onePathSum(root.Left, sum - root.Val, onePath)
    onePathSum(root.Right, sum - root.Val, onePath)

    onePath = onePath[:len(onePath) - 1]
}
