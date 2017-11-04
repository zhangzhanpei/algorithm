/**
 * 给定一个旋转有序数组和一个目标值, 判断目标值是否在数组中
 */
func search(nums []int, target int) bool {
    for _, v := range nums {
        if v == target {
            return true
        }
    }
    return false
}
