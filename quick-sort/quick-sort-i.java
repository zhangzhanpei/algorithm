/**
 * 三路快排
 * 将数组划分成 < pivot, = pivot, > pivot 三部分
 */
public static void quickSort(int[] nums, int left, int right) {
    if (left >= right) {
        return;
    }
    
    // 取第一个元素为比较基数
    int pivot = nums[left];
    int i = left;
    int k = left + 1;
    int j = right;

    OUTLOOP:
    while (k < j) {
        if (nums[k] < pivot) {
            swap(nums, i, k);
            i++;
            k++;
        } else if (nums[k] == pivot) {
            k++;
        } else { // nums[k] > pivot 时
            while (nums[j] > pivot) {
                j--;
                if (j < k) { // 如果右边都是大于 pivot 的元素，说明右边已经不用处理了
                    break OUTLOOP;
                }
            }
            if (nums[j] == pivot) {
                swap(nums, k, j);
                k++;
                j--;
            }
            if (nums[j] < pivot) {
                swap(nums, i, j);
                swap(nums, j, k);
                i++;
                k++;
                j--;
            }
        }
    }

    quickSort(nums, left, i - 1);
    quickSort(nums, j + 1, right);
}

/**
 * 交换数组元素位置
 */
public static void swap(int[] nums, int i, int j) {
    if (i == j) {
        return;
    }
    int tmp = nums[i];
    nums[i] = nums[j];
    nums[j] = tmp;
}
