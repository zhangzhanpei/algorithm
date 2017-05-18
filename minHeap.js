//最小堆, 刚好和最大堆相反, 堆顶就是最小的元素
//交换数组元素
function swap(arr, i, j)
{
    let tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

//维护最小堆
function minHeap(arr, i, len)
{
    //左右儿子下标
    let l = i * 2 + 1;
    let r = i * 2 + 2;
    let min = i;
    //当前节点与左右儿子比较, 得到最小的节点下标
    if (l < len && arr[l] < arr[i]) {
        min = l;
    }
    if (r < len && arr[r] < arr[min]) {
        min = r;
    }

    //交换最小节点
    if (min != i) {
        swap(arr, min, i);
        //此时min下标的值交换为原来i结点的大值, 需要继续对min下标节点进行最小堆维护
        minHeap(arr, min, len);
    }
}

//对随机数组建立最小堆
function buildMinHeap(arr)
{
    //最小堆中, arr.length / 2 到 arr.length - 1 都是叶子节点, 所以只需对前面一半的元素进行最小堆调整
    for (let i = Math.floor(arr.length / 2) + 1; i >= 0; i--) {
        minHeap(arr, i, arr.length);
    }
}

//最小堆排序
function minHeapSort(arr)
{
    //对乱序数组建立最小堆
    buildMinHeap(arr);
    //此时最小堆顶部元素[即数组第一个元素]就是数组最小值
    let len = arr.length - 1;
    //当数组长度<=1时不用再处理了
    while (len > 1) {
        //将arr[0]元素交换到数组尾部
        swap(arr, 0, len);
        //交换后 arr 的 0 到 len-1 位置的元素可能不再符合最小堆, 对 arr[0] 进行维护
        len--;
        minHeap(arr, 0, len);
    }
}

let a = [10, 8, 11, 8, 14, 9 , 4, 1, 17];
minHeapSort(a);
console.log(a);
