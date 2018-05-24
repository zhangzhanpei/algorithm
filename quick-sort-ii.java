/**
 * 双轴快排
 */
public static void quickSort(int[] nums, int left, int right) {
    if (left >= right) {
        return;
    }
    // 保证 pivotLeft < pivotRight
    if (nums[left] > nums[right]) {
        swap(nums, left, right);
    }
    // 取左右两端元素为比较基数
    int pivotLeft = nums[left];
    int pivotRight = nums[right];
    int i = left;
    int k = left + 1;
    int j = right;

    OUTLOOP:
    while (k < j) {
        if (nums[k] < pivotLeft) {
            i++;
            swap(nums, i, k);
            k++;
        } else if (nums[k] >= pivotLeft && nums[k] <= pivotRight) {
            k++;
        } else { // nums[k] > pivotRight
            // 从 nums[j - 1] 向左扫，找到第一个 <= pivotRight 的元素，再分情况处理
            j--;
            while (nums[j] > pivotRight) {
                // 当 j 和 k 相遇即扫完一趟
                if (j <= k) {
                    break OUTLOOP;
                }
                j--;
            }
            // 交换元素到 >= pivotLeft && <= pivotRight 区间
            if (nums[j] >= pivotLeft && nums[j] <= pivotRight) {
                swap(nums, j, k);
                k++;
            } else { // nums[j] < pivotLeft，交换元素到 < pivotLeft 区间
                i++;
                swap(nums, j, k);
                swap(nums, i, k);
                k++;
            }
        }
    }

    // 将两个比较基数交换到正确位置
    // nums[left + 1] 到 nums[i] 的元素都小于 pivotLeft，将 nums[left] 和 nums[i] 交换后，nums[left] 到 nums[i - 1] 都小于 pivotLeft
    swap(nums, left, i);
    // nums[j] 到 nums[right - 1] 的元素都大于 pivotRight，将 nums[right] 和 nums[j] 交换后，nums[j + 1] 到 nums[right] 都大于 pivotRight
    swap(nums, right, j);

    // 双轴将数组分成三部分
    quickSort(nums, left, i - 1); // 递归处理小于 pivotLeft 部分
    quickSort(nums, i + 1, j - 1); // 递归处理大于 pivotLeft 小于 pivotRight 部分
    quickSort(nums, j + 1, right); // 递归处理大于 pivotRight 部分
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
