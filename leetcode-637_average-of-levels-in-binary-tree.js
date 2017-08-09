//求二叉树每一层的平均值
//BFS广度优先遍历
var averageOfLevels = function(root) {
    let list = [], ret = [];
    if (root === null) {
        return ret;
    }
    list.push(root); //list一开始先保存第一层的元素root节点
    while (list.length) {
        let len = list.length; //获取当前层的宽度
        let sum = 0;
        for (let i = 0; i < len; i++) { //遍历当前层时, 将下一层的元素压进list尾部
            let node = list.shift(); //当前层的元素逐个弹出, 当前层遍历完后剩下的元素就全是下一层的
            sum += node.val;
            if (node.left !== null) {
                list.push(node.left);
            }
            if (node.right !== null) {
                list.push(node.right);
            }
        }
        ret.push(sum / len);
    }
    return ret;
};

