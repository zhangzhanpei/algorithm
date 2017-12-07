/**
 * 快速排序
 */
public int[] quickSort(int left, int right, int[] arr) {
    if (left > right) {
        return null;
    }
    int i = left;
    int j = right;
    int base = arr[left]; //取第一个元素作为比较
    while (i != j) {
        while (arr[j] >= base && i < j) { //从右往左找到第一个小于 base 的元素
            j--;
        }
        while (arr[i] <= base && i < j) { //从左往右找到第一个大于 base 的元素
            i++;
        }
        if (i < j) { //两个元素在基数左右两边时交换两数
            int tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
        }
    }

    //一趟之后，比基数小的交换到了左边，比基数大的交换到了右边
    arr[left] = arr[i]; //将基数交换到 i 和 j 相遇的位置
    arr[i] = base;

    //递归继续处理
    quickSort(left, i - 1, arr);
    quickSort(i + 1, right, arr);
    return arr;
}
