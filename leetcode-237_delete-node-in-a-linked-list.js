//删除单链表中的某个节点
var deleteNode = function(node) {
    //删除node时, 将node的值设为下一个元素的值, 并将next设为next.next, 即跳过[删除]当前node节点
    node.val = node.next.val;
    node.next = node.next.next;
};
