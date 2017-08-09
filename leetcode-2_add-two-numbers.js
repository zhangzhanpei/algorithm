//给定两个单链表形式的数字, 低位在前, 如 2->4->3 表示 342, 5->6->4 表示 465
//两数相加得到 807, 表示为 7->0->8
//两个链表同时跑, 逐个相加直到两个链表的元素都跑完

function addTwoNumbers(l1, l2) {
    let sumLinkList = new ListNode(-1); //定义一个结果链表
    let head = sumLinkList;
    let carry = 0; //进位
    while (l1 || l2) { //两个链表相同位置的元素相加, 如果其中一个链表跑完了, 则用0代替
        let a = 0;
        let b = 0;
        if (l1) {
            a = l1.val;
            l1 = l1.next;
        } else {
            a = 0;
        }

        if (l2) {
            b = l2.val;
            l2 = l2.next;
        } else {
            b = 0;
        }

        let sum = a + b + carry; //两个元素相加, 加上上一次的进位
        if (sum >= 10) {
            carry = 1;
            sumLinkList.next = new ListNode(sum - 10);
        } else {
            carry = 0;
            sumLinkList.next = new ListNode(sum);
        }
        sumLinkList = sumLinkList.next;
    }

    if (carry === 1) {
        sumLinkList.next = new ListNode(1);
    }

    return head.next;
}

let sumLinkList = addTwoNumbers(l11, l21);
while (sumLinkList) {
    console.log(sumLinkList.val);
    sumLinkList = sumLinkList.next;
}
