/**
 * 给定一个旋转有序数组(包含重复元素), 找到最小元素
 */
func findMin(nums []int) int {
    ret := nums[0]
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i + 1] {
            ret = nums[i + 1]
        }
    }
    return ret
}
