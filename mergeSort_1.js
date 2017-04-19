//自底向上的归并排序
//合并两个有序数组
function mergeArray(arr, left, mid, right)
{
    let i = left;
    let j = mid + 1;
    let k = 0; //临时存放合并结果数组的下标
    let tmp = []; //临时数组

    //合并两个有序数组直到其中一个数组先取完
    while (i <= mid && j <= right) {
        if (arr[i] <= arr[j]) {
            tmp[k] = arr[i]
            i++;
            k++;
        } else {
            tmp[k] = arr[j];
            j++;
            k++;
        }
    }

    //当右侧数组读完后，把左侧数组剩下部分追加到临时数组
    while (i <= mid) {
        tmp[k] = arr[i];
        i++;
        k++;
    }

    //当左侧数组读完后，把右侧数组剩下部分追加到临时数组
    while (j <= right) {
        tmp[k] = arr[j];
        j++;
        k++;
    }

    //将临时合并有序的数组填回原数组，因为要用原数组进行元素比较
    for (k = 0, i = left; i <= right; i++, k++) {
        arr[i] = tmp[k];
    }
}

//一一合并，两两合并，四四合并，八八合并...
function mergeSort(arr)
{
    let len = arr.length;
    let mergeSize = 1;
    while (mergeSize < len) {
        console.log(mergeSize);
        let left = mid = right = 0;
        while (left + mergeSize * 2 - 1 < len) { //未合并的元素要大于 mergeSize * 2 才足够继续合并
            mid = left + mergeSize -1; //因为是左右两个 mergeSize 的元素进行合并，所以 mid 的下标为 left + mergeSize - 1
            right = left + mergeSize * 2 -1; //同理可得 right = left + mergeSize * 2 - 1
            mergeArray(arr, left, mid, right);
            left = right + 1; //向右移动，处理下两个待合并的元素数组
        }
        //当 meregeSize 增长到整个数组无法拆分成两个 mergeSize 大小的数组时，拆为左边一个 mergeSize 大小的数组，右边则是剩下的元素，此时 mid = left + mergeSize - 1
        mid = left + mergeSize - 1;
        if (mid < len - 1) {
            mergeArray(arr, left, mid, len - 1); //直接合并最后一次数组
        }
        mergeSize *= 2;
    }
    return arr;
}

var testArr = [6, 2, 13, 10, 1, 0, 89, 24, 11];
mergeSort(testArr);
