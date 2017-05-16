//最大堆
//交换数组元素
function swap(arr, i, j)
{
    let tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

//维护最大堆
function maxHeap(arr, i, len)
{
    //左右儿子下标
    let l = i * 2 + 1;
    let r = i * 2 + 2;
    let max = i;
    //当前节点与左右儿子比较, 得到最大的节点下标
    if (l < len && arr[l] > arr[i]) {
        max = l;
    }
    if (r < len && arr[r] > arr[max]) {
        max = r;
    }

    //交换最大节点
    if (max != i) {
        swap(arr, max, i);
        //此时max下标的值交换为原来i结点的小值, 需要继续对小节点进行最大堆维护
        maxHeap(arr, max, len);
    }
}

//对随机数组建立最大堆
function buildMaxHeap(arr)
{
    //最大堆中, arr.length / 2 到 arr.length - 1 都是叶子节点, 所以只需对前面一半的元素进行最大堆调整
    for (let i = Math.floor(arr.length / 2) + 1; i >= 0; i--) {
        maxHeap(arr, i, arr.length);
    }
}

//最大堆排序
function maxHeapSort(arr)
{
    //对乱序数组建立最大堆
    buildMaxHeap(arr);
    //此时最大堆顶部元素[即数组第一个元素]就是数组最大值
    let len = arr.length - 1;
    //当数组长度<=1时不用再处理了
    while (len > 1) {
        //将arr[0]元素交换到数组尾部
        swap(arr, 0, len);
        //交换后 arr 的 0 到 len-1 位置的元素可能不再符合最大堆, 对 arr[0] 进行维护
        len--;
        maxHeap(arr, 0, len);
    }
}

let a = [10, 8, 11, 8, 14, 9 , 4, 1, 17];
maxHeapSort(a);
console.log(a);
