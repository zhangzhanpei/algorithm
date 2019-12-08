/**
 * 给定一个单链表, 返回随机一个节点[保证取得每个节点的概率一样]
 * 解一: 遍历整个链表取得长度再取随机
 * 解二: 水塘抽样法
 * 取第一个节点时概率为 1/1
 * 取第二个节点时概率为 1/2
 * ...
 * 取第 n 个节点时概率为 1/n
 */
class Solution {

    ListNode head;

    public Solution(ListNode head) {
        this.head = head;
    }

    public int getRandom() {
        int i = 0;
        ListNode current = this.head;
        int ret = current.val;
        while (current != null) {
            i++;
            Random random = new Random();
            if (random.nextInt(i) == 0) { // 当来到 current 节点时长度为 i, 从 0-i 中取到 0 的概率为 1/i, 此时可以取 current 的值
                ret = current.val;
            }
            current = current.next;
        }
        return ret;
    }
}
