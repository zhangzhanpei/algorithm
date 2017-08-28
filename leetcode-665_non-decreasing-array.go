/**
 * 给定一个int数组, 判断数组是否是增序的(允许调整其中一个元素的大小)
 */
package main

func main() {
    checkPossibility([]int{2, 4, 2, 3})
}

func checkPossibility(nums []int) bool {
    count := 0
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i + 1] {
            count++
            if i > 0 && nums[i + 1] < nums[i - 1] { //如果下一个数比前一个数小
                nums[i + 1] = nums[i] //将下一个数赋值为当前数
            } else { //如果下一个数比前一个数大
                nums[i] = nums[i + 1] //将当前数赋值为下一个数
            }
        }
    }
    return count <= 1
}
