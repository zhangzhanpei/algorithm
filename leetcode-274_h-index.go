/**
 * 给定一个数组表示论文引用次数, 求H指数
 * 将引用次数由高到低排序, 从前往后查有序数组, 如果下标大于元素值, 那么该下标就是H指数
 */
func hIndex(citations []int) int {
    citations = sort(citations)
    for i := 0; i < len(citations); i++ {
        if i >= citations[i] {
            return i
        }
    }
    return len(citations)
}

//由大到小排序数组
func sort(nums []int) []int {
    for i := 1; i < len(nums); i++ { //遍历数组元素
        for j := i; j >= 1 && nums[j] > nums[j - 1]; j-- { //如果当前元素比前面的小, 一直向前交换直到位置正确
            nums[j - 1], nums[j] = nums[j], nums[j - 1]
        }
    }
    return nums
}
