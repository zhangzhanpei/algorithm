//自顶向下的归并排序
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

//递归拆分数组
function mergeSort(arr, left, right)
{
    //递归调用mergeSort进行数组拆分，当left=right时说明拆到了最小的数组，只剩一个元素
    //当无法继续往下递归拆分数组时，返回上一层调用，开始拆分上一层右侧的数组
    //当left=0,right=0时,回到上一层left=0,right=1执行mergeSort(arr, 1, 1)拆右边，此时右边也拆完，则执行mergeArray(arr, 0, 1, 1)进行合并数组
    if (left < right) {
        let mid = parseInt((left + right) / 2);
        mergeSort(arr, left, mid);
        mergeSort(arr, mid + 1, right);
        mergeArray(arr, left, mid, right);
    }
    return arr;
}

var testArr = [6, 2, 13, 10, 1, 0, 89, 24, 11, 100];
mergeSort(testArr, 0, testArr.length - 1);
