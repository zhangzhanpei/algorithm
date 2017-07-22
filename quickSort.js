//快速排序
function quickSort(left, right, arr) {
    if (left > right) {
        return;
    }
    var i = left; //两个指针
    var j = right;
    var base = arr[left]; //基数
    while (i != j) {
        while (arr[j] >= base && i < j) { //从右往左找到第一个比基数小的数
            j--;
        }
        while (arr[i] <= base && i < j) { //从左往右找到第一个比基数大的数
            i++;
        }
        if (i < j) { //如果i和j没有相遇, 则交换两数
            let tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
        }
    }
    arr[left] = arr[i]; //i和j相遇后, 即一趟过后交换基数和相遇位置的数, 此时左侧的数都比基数小, 右侧的数都比基数大
    arr[i] = base;

    quickSort(left, i - 1, arr); //递归处理左右两半数据
    quickSort(i + 1, right, arr);
}

var arr = [3, 8, 1, 16, 5, 0, 20];
quickSort(0, arr.length - 1, arr);
