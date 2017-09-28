/**
 * 给定一个旋转有序数组和一个目标值, 返回目标值在数组中的下标
 */
func search(nums []int, target int) int {
    for k, v := range nums {
        if v == target {
            return k
        }
    }
    return -1
}
