//求数组中第k大的数
package main

import "fmt"

func main() {
    fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
    low := 0
    high := len(nums)
    for low < high { //使用快排, 比基数小的交换到左边, 大的交换到右边
        i := low
        j := high - 1
        base := nums[low]
        for i != j {
            for i < j && nums[i] >= base {
                i++
            }
            for i < j && nums[j] <= base {
                j--
            }
            if i < j {
                nums[i], nums[j] = nums[j], nums[i] //交换基数和i, j相遇处的元素
                i++
                j--
            }

            if j == k - 1 { //判断 i 和 j 相遇处是否为第 k 位
                return nums[j]
            } else if j < k - 1 {
                low = j + 1
            } else {
                high = j
            }
        }
    }
    return 0
}
