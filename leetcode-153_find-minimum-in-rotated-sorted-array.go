/**
 * 给定一个旋转有序数组, 找到最小元素
 */

//O(n)复杂度
func findMin(nums []int) int {
    ret := nums[0]
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i + 1] {
            ret = nums[i + 1]
        }
    }
    return ret
}

//O(lgn)复杂度
func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    if nums[left] > nums[right] { //旋转有序数组, 如果开头元素大于结尾元素, 说明最小值在中间
        for left + 1 != right {
            mid := (left + right) / 2
            if nums[left] < nums[mid] { //如果开头元素小于中点元素, 说明中点左侧都是有序的, 最小值在右侧
                left = mid
            } else {
                right = mid
            }
        }
        return int(math.Min(float64(nums[left]), float64(nums[right])))
    }
    return nums[0] //如果开头元素小于结尾元素, 那么开头元素就是最小的
}
