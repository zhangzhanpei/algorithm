//给定一个单链表, 判断它是否有环, 如果有则返回环的起点, 如果不是返回null
var detectCycle = function(head) {
    if (head === null || head.next === null) {
        return null;
    }
    var slow = head;
    var fast = head;
    var entry = head;
    while (fast.next && fast.next.next) {
        slow = slow.next;
        fast = fast.next.next;
        if (slow === fast) { //快慢指针相遇, 此时开始找环的起点
            while (slow != entry) {
                slow = slow.next;
                entry = entry.next;
            }
            return entry;
        }
    }
    return null;
};

/**
 * 用2个辅助指针，第一个指针p1每次走一步，第二个指针p2每次走两步。
 * 当p1走到环的起点时走了a步，那么p2肯定在环中，而且走了2a步，距离环的起点是2a-a步。
 * 如果p1再走b步与p2相遇，那么就会有这么一个等式
 * a + b = 2 * (a + b) - k * n
 *
 * 其中，k就是b在环中走的圈数，n是环的长度。因为相遇时可能p2在环中走了k圈了。上述等式可以转换为下式
 * a + b = k * n
 *
 * 这时，把p1拉回链表起点，p2留在相遇点。
 * p1和p2每次都走一步，那么当p1走了a步后，也就是环的起点时，p2也走了a步。
 * 但是因为p2已经在环的b位置，那么再走a步后，根据上面的公式，就会发现p2走到了环的起点。
 * 因为p1也在环的起点，所以2个指针就会相遇，而相遇点就是环的起点。
 */
