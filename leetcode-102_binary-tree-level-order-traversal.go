/**
 * 按层遍历二叉树, 每层节点以数组形式返回
 * BFS广度优先遍历
 */
func levelOrder(root *TreeNode) [][]int {
    list, ret := []*TreeNode{}, [][]int{}
    if root == nil {
        return ret
    }
    list = append(list, root)
    for len(list) > 0 {
        n := len(list)
        tmp := []int{}
        for i := 0; i < n; i++ {
            tmp = append(tmp, list[0].Val)
            if list[0].Left != nil {
                list = append(list, list[0].Left)
            }
            if list[0].Right != nil {
                list = append(list, list[0].Right)
            }
            list = list[1:] //第一个元素弹出
        }
        ret = append(ret, tmp)
    }
    return ret
}
