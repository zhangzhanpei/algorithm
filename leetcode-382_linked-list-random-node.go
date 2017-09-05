/**
 * 给定一个单链表, 返回随机一个节点[保证取得每个节点的概率一样]
 * 解一: 遍历整个链表取得长度再取随机
 * 解二: 水塘抽样法
 * 取第一个节点时概率为 1/1
 * 取第二个节点时概率为 1/2
 * ...
 * 取第 n 个节点时概率为 1/n
 */
type Solution struct {
    Head *ListNode
}

func Constructor(head *ListNode) Solution {
    return Solution{Head : head}
}

func (this *Solution) GetRandom() int {
    i := 0
    cur := this.Head
    ret := cur.Val
    for cur != nil {
        i++
        if rand.Intn(i) == 0 { //当来到 cur 节点时长度为 i, 从 0-i 中取到 0 的概率为 1/i, 此时可以取 cur 的值 
            ret = cur.Val
        }
        cur = cur.Next
    }
    return ret
}
