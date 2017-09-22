/**
 * 同 leetcode-274, 数组改为有序数组, 此时可以改用二分查找
 */
func hIndex(citations []int) int {
    n := len(citations)
    left, right := 0, n - 1
    for left <= right {
        mid := (left + right) / 2
        if citations[mid] == n - mid {
            return n - mid
        } else if citations[mid] > n - mid {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return n - left
}
